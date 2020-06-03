package gql

import (
	"github.com/adyatlov/xp/plugin"
)

type pluginResolver struct {
	plugin plugin.Plugin
}

func (r *pluginResolver) Name() string {
	return string(r.plugin.Name())
}

func (r *pluginResolver) Description() string {
	return r.plugin.Description()
}
