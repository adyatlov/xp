package node

import (
	"fmt"

	"github.com/adyatlov/bunxp/objects"
	"github.com/mesosphere/bun/v2/bundle"
)

var metricNodeType = objects.MetricType{
	ValueType:      objects.MTType,
	ObjectTypeName: "node",
	Name:           "node-type",
	MetricName:     "Node Type",
	Description:    "Type of the DC/OS node; can be \"master\", \"agent\" or \"public agent\"",
	Evaluate:       metricNodeTypeEvaluate,
}

func metricNodeTypeEvaluate(b *bundle.Bundle, id objects.ObjectId) (interface{}, error) {
	host, ok := b.Hosts[string(id)]
	if !ok {
		return nil, fmt.Errorf("cannot find a node with id \"%v\" in the bundle", id)
	}
	return host.Type, nil
}
