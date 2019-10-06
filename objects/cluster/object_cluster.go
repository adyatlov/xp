package cluster

import (
	"fmt"

	"github.com/adyatlov/bunxp/objects"

	"github.com/mesosphere/bun/v2/bundle"
)

func init() {
	cluster := objects.ObjectType{
		Name:              "cluster",
		DisplayName:       "Cluster",
		PluralDisplayName: "Clusters",
		Description:       "DC/OS Cluster",
		Find:              find,
		DefaultMetrics:    []objects.MetricTypeName{"dcos-version"},
		Metrics: []objects.MetricType{
			metricVersion,
		},
	}
	objects.RegisterObjectType(cluster)
}

func find(b *bundle.Bundle, id objects.ObjectId, withChildren bool) (*objects.Object, error) {
	object := &objects.Object{}
	setNameAndId(b, object)
	if withChildren {
		findNodes(b, object)
	}
	return object, nil
}

func setNameAndId(b *bundle.Bundle, object *objects.Object) {
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
	object.Id = objects.ObjectId(config.ClusterName)
	object.Name = objects.ObjectName(config.ClusterName)
	if object.Id == "" {
		object.Id = "*Unknown ID*"
		object.Name = "*Unknown Name*"
	}
}

func findNodes(b *bundle.Bundle, o *objects.Object) {
	nodeType := objects.MustGetObjectType("node")
	nodes := objects.Children{
		Type:    nodeType.Name,
		Objects: make([]*objects.Object, 0, len(b.Hosts)),
	}
	for ip, _ := range b.Hosts {
		obj, err := nodeType.New(b, objects.ObjectId(ip), false)
		if err != nil {
			o.Errors = append(o.Errors, fmt.Sprintf("Cannot create node \"%s\": %s", ip, err.Error()))
		}
		nodes.Objects = append(nodes.Objects, obj)
	}
	o.Children = append(o.Children, nodes)
	return
}
