package gql

import (
	"github.com/adyatlov/xp/data"
	"github.com/graph-gophers/graphql-go"
)

type pluginResolver struct {
	plugin *data.Plugin
}

func (r pluginResolver) Id() graphql.ID {
	return encodeId(pluginId{PluginName: r.plugin.Name})
}

func (r pluginResolver) Name() string {
	return string(r.plugin.Name)
}

func (r pluginResolver) Description() string {
	return r.plugin.Description
}
