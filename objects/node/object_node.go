package node

import (
	"fmt"

	"github.com/adyatlov/bunxp/objects"

	"github.com/mesosphere/bun/v2/bundle"
)

func init() {
	node := &objects.ObjectType{
		Name:           "node",
		Description:    "DC/OS Node",
		Find:           find,
		DefaultMetrics: []objects.MetricTypeName{"node-type"},
		Metrics: []*objects.MetricType{
			metricNodeType,
		},
	}
	objects.RegisterObjectType(node)
}

func find(b *bundle.Bundle, id objects.ObjectId) (*objects.Object, error) {
	_, ok := b.Hosts[string(id)]
	if !ok {
		return nil, fmt.Errorf("cannot find a node with id \"%v\" in the bundle", id)
	}
	return &objects.Object{
		Name: objects.ObjectName(id),
	}, nil
}
