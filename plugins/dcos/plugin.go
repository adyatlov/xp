package dcos

import (
	"github.com/adyatlov/xp/data"
)

func init() {
	p := &data.Plugin{
		Name:        "DC/OS Cluster",
		Description: "Plugin for DC/OS diagnostics bundle",
		Open:        open,
		Compatible:  compatible,
	}
	data.RegisterPlugin(p)
}

func open(string) (data.Dataset, error) {
	return nil, nil
}

func compatible(url string) (bool, error) {
	if url == "example.com/?minEmployee=10&maxEmployee=100&nDivision=11" {
		return true, nil
	}
	return false, nil
}
