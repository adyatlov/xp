package dcos

import (
	"github.com/adyatlov/xp/data"
	"github.com/adyatlov/xp/plugin"
)

func init() {
	p := &Plugin{}
	plugin.RegisterPlugin(p)
}

type Plugin struct {
}

func (p Plugin) Name() plugin.Name {
	return "DC/OS Cluster"
}

func (p Plugin) Description() string {
	return "Plugin for DC/OS diagnostics bundle"
}

func (p Plugin) Open(url string) (data.Dataset, error) {
	return nil, nil
}

func (p Plugin) Compatible(url string) (bool, error) {
	if url == "example.com/?minEmployee=10&maxEmployee=100&nDivision=11" {
		return true, nil
	}
	return false, nil
}
