package dcos

import (
	"github.com/adyatlov/xp/data"
	"github.com/mesosphere/bun/v2/bundle"
)

type agent struct {
	ip       string
	id       string
	t        string
	hostname interface{}
}

func (a agent) Type() *data.ObjectType {
	return TAgent
}

func (a agent) Id() data.ObjectId {
	return data.ObjectId(a.ip)
}

func (a agent) Name() data.ObjectName {
	return data.ObjectName(a.ip)
}

func (a agent) Properties(properties *[]interface{}, names ...data.PropertyName) error {
	for _, name := range names {
		switch name {
		case PId.Name:
			*properties = append(*properties, a.id)
		case PAgentType.Name:
			*properties = append(*properties, a.t)
		case PAgentHostname.Name:
			*properties = append(*properties, a.hostname)
		default:
			*properties = append(*properties, nil)
		}
	}
	return nil
}

func (a agent) ChildGroup(childTypeName data.ObjectTypeName) data.ObjectGroup {
	return nil
}

func newAgent(host bundle.Host) *agent {
	a := agent{}
	a.ip = string(host.IP)
	id := struct {
		Id string `json:"id"`
	}{}
	err := host.ReadJSON("mesos-agent-state", &id)
	if err != nil {
		a.id = "N/A"
	} else {
		a.id = id.Id
	}
	a.t = string(host.Type)
	hostname := struct {
		Hostname string `json:"hostname"`
	}{}
	err = host.ReadJSON("diagnostics-health", &hostname)
	if err != nil {
		a.hostname = "N/A"
	} else {
		a.hostname = hostname.Hostname
	}
	return &a
}

type agentGroup struct {
	b          bundle.Bundle
	totalCount int
}

func newAgentGroup(b bundle.Bundle) *agentGroup {
	return &agentGroup{
		b:          b,
		totalCount: len(b.Agents()),
	}
}

func (a agentGroup) Type() *data.ObjectType {
	return TAgent
}

func (a agentGroup) All(agents *[]data.Object) error {
	hosts := a.b.Agents()
	hosts = append(hosts, a.b.PublicAgents()...)
	for _, host := range hosts {
		*agents = append(*agents, newAgent(host))
	}
	return nil
}

func (a agentGroup) TotalCount() int {
	return a.totalCount
}

func (a agentGroup) Pager() data.ObjectPager {
	return nil
}
