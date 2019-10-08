package agenttype

import (
	"fmt"

	"github.com/adyatlov/bunxp/explorer"

	"github.com/mesosphere/bun/v2/bundle"
)

func init() {
	t := explorer.MetricType{
		ValueType:   explorer.MTType,
		Name:        "agent-type",
		MetricName:  "Agent Type",
		Description: "Type of the DC/OS node; can be \"agent\" or \"public agent\"",
		Evaluate:    e,
	}
	explorer.RegisterMetricType(t)
}

func e(b *bundle.Bundle, id explorer.ObjectId) (interface{}, error) {
	host, ok := b.Hosts[string(id)]
	if !ok {
		return nil, fmt.Errorf("cannot find an agent with id \"%v\" in the bundle", id)
	}
	return host.Type, nil
}
