package cluster

import (
	"fmt"

	"github.com/adyatlov/bunxp/explorer"

	"github.com/mesosphere/bun/v2/bundle"
)

func init() {
	t := explorer.ObjectType{
		Name:              "cluster",
		DisplayName:       "Cluster",
		PluralDisplayName: "Clusters",
		Description:       "DC/OS Cluster",
		Find:              find,
		DefaultMetrics:    []explorer.MetricTypeName{"dcos-version"},
		Metrics:           []explorer.MetricTypeName{"dcos-version"},
	}
	explorer.RegisterObjectType(t)
}

func find(b *bundle.Bundle, id explorer.ObjectId, withChildren bool) (*explorer.Object, error) {
	object := &explorer.Object{}
	setNameAndId(b, object)
	if withChildren {
		findAgent(b, object)
	}
	return object, nil
}

func setNameAndId(b *bundle.Bundle, object *explorer.Object) {
	if len(b.Hosts) == 0 {
		return
	}
	var host bundle.Host
	for _, host = range b.Hosts {
		break
	}
	config := &struct {
		ClusterName string `json:"cluster_name"`
	}{}
	if err := host.ReadJSON("expanded-config", config); err != nil {
		object.Errors = append(object.Errors,
			fmt.Sprintf("cannot parse cluster name: %s", err.Error()))
	}
	object.Id = explorer.ObjectId(config.ClusterName)
	object.Name = explorer.ObjectName(config.ClusterName)
	if object.Id == "" {
		object.Id = "*Unknown ID*"
		object.Name = "*Unknown Name*"
	}
}

func findAgent(b *bundle.Bundle, o *explorer.Object) {
	t := explorer.MustGetObjectType("agent")
	nodes := explorer.ObjectGroup{
		Type:    t.Name,
		Objects: make([]*explorer.Object, 0, len(b.Hosts)),
	}
	for ip, _ := range b.Hosts {
		obj, err := t.New(b, explorer.ObjectId(ip), false)
		if err != nil {
			o.Errors = append(o.Errors, fmt.Sprintf("Cannot create agent \"%s\": %s", ip, err.Error()))
		} else {
			nodes.Objects = append(nodes.Objects, obj)
		}
	}
	o.Children = append(o.Children, nodes)
	return
}
