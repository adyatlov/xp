package gql

import "github.com/adyatlov/xp/data"

type pluginResolver struct {
	plugin data.Plugin
}

func (r *pluginResolver) Name() string {
	return string(r.plugin.Name())
}

func (r *pluginResolver) Description() string {
	return r.plugin.Description()
}
