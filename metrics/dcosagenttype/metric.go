package dcosagenttype

import (
	"fmt"

	"github.com/adyatlov/xp/xp"

	"github.com/mesosphere/bun/v2/bundle"
)

func init() {
	t := xp.MetricType{
		ValueType:   xp.MTType,
		Name:        "dcos-agent-type",
		DisplayName: "DC/OS Agent Type",
		Description: "Type of the DC/OS agent; can be \"agent\" or \"public agent\"",
		Evaluate:    e,
	}
	xp.RegisterMetricType(t)
}

func e(b *bundle.Bundle, id xp.ObjectId) (interface{}, error) {
	host, ok := b.Hosts[string(id)]
	if !ok {
		return nil, fmt.Errorf("cannot find an agent with id \"%v\" in the bundle", id)
	}
	return host.Type, nil
}
