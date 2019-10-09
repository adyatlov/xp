package dcosagent

import (
	"fmt"

	"github.com/adyatlov/bunxp/explorer"

	"github.com/mesosphere/bun/v2/bundle"
)

func init() {
	t := explorer.ObjectType{
		Name:              "agent",
		DisplayName:       "DC/OS Agent",
		PluralDisplayName: "DC/OS Agents",
		Description:       "DC/OS Agent is a DC/OS worker node",
		Find:              find,
		Metrics:           []explorer.MetricTypeName{"dcos-agent-type"},
		DefaultMetrics:    []explorer.MetricTypeName{"dcos-agent-type"},
	}
	explorer.RegisterObjectType(t)
}

func find(b *bundle.Bundle, id explorer.ObjectId, withChildren bool) (*explorer.Object, error) {
	host, ok := b.Hosts[string(id)]
	if !ok || !(host.Type == bundle.DTAgent || host.Type == bundle.DTPublicAgent) {
		return nil, fmt.Errorf("cannot find an agent with id \"%v\" in the bundle", id)
	}
	return &explorer.Object{
		Name: explorer.ObjectName(id),
	}, nil
}
