package dcosagent

import (
	"fmt"

	"github.com/adyatlov/bunxp/xp"

	"github.com/mesosphere/bun/v2/bundle"
)

func init() {
	t := xp.ObjectType{
		Name:              "agent",
		DisplayName:       "DC/OS Agent",
		PluralDisplayName: "DC/OS Agents",
		Description:       "DC/OS Agent is a DC/OS worker node",
		Metrics:           []xp.MetricTypeName{"dcos-agent-type"},
		DefaultMetrics:    []xp.MetricTypeName{"dcos-agent-type"},
		FindObject:        findObject,
		FindChildren:      nil,
	}
	xp.RegisterObjectType(t)
}
func findObject(b *bundle.Bundle, id xp.ObjectId) (xp.ObjectName, error) {
	host, ok := b.Hosts[string(id)]
	if !ok || !(host.Type == bundle.DTAgent || host.Type == bundle.DTPublicAgent) {
		return "", fmt.Errorf("cannot find an agent with id \"%v\" in the bundle", id)
	}
	return xp.ObjectName(id), nil
}
