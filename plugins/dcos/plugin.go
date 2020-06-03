package dcos

import (
	"time"

	"github.com/adyatlov/xp/data"
)

func init() {
	plugin := &Plugin{}
	data.RegisterPlugin(plugin)
}

type Plugin struct {
}

func (p Plugin) Name() data.PluginName {
	return "DC/OS Cluster"
}

func (p Plugin) Description() string {
	return "Plugin for DC/OS diagnostics bundle"
}

func (p Plugin) Open(url string) (data.Dataset, error) {
	return nil, nil
}

func (p Plugin) Compatible(url string) (bool, error) {
	time.Sleep(1 * time.Second)
	if url == "example.com/?minEmployee=10&maxEmployee=100&nDivision=11" {
		return true, nil
	}
	return false, nil
}
