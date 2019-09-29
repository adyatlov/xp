package cluster

import (
	"fmt"

	"github.com/adyatlov/bunxp/objects"

	"github.com/mesosphere/bun/v2/bundle"
)

func init() {
	cluster := objects.ObjectType{
		Name:           "cluster",
		Description:    "DC/OS Cluster",
		Find:           find,
		Parents:        parents,
		Children:       children,
		DefaultMetrics: []objects.MetricName{"version"},
		Metrics: []objects.Metric{
			metricVersion(),
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
		object.Name = "*Unknown Name*"
	}
	return object, nil
}

func parents(b *bundle.Bundle, id objects.ObjectId) ([]*objects.Object, error) {
	return nil, nil
}

func children(b *bundle.Bundle, id objects.ObjectId) ([]*objects.Object, error) {
	return nil, nil
}

func metricVersion() objects.Metric {
	metric := objects.Metric{
		Type:        objects.MTVersion,
		Name:        "version",
		Description: "DC/OS Version installed on the cluster",
	}
	metric.Evaluate = func(b *bundle.Bundle, o *objects.Object) (interface{}, error) {
		if len(b.Hosts) == 0 {
			return nil, fmt.Errorf("cannot find a single host in the directory %s", b.Path)
		}
		var host bundle.Host
		for _, host = range b.Hosts {
			break
		}
		version := &struct {
			Version string
		}{}
		if err := host.ReadJSON("dcos-version", version); err != nil {
			return nil, err
		}
		return version.Version, nil
	}
	return metric
}
