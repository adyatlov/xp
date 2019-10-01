package cluster

import (
	"fmt"

	"github.com/adyatlov/bunxp/objects"

	"github.com/mesosphere/bun/v2/bundle"
)

func init() {
	cluster := &objects.ObjectType{
		Name:           "cluster",
		Description:    "DC/OS Cluster",
		Find:           find,
		DefaultMetrics: []objects.MetricTypeName{"dcos-version"},
		Metrics: []*objects.MetricType{
			metricVersion,
		},
	}
	objects.RegisterObjectType(cluster)
}

func find(b *bundle.Bundle, id objects.ObjectId) (*objects.Object, error) {
	object := &objects.Object{}
	if len(b.Hosts) != 0 {
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
	}
	if object.Id == "" {
		object.Id = "*Unknown ID*"
		object.Name = "*Unknown MetricName*"
	}
	return object, nil
}
