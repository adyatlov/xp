package gql

import (
	"github.com/adyatlov/xp/plugin"
	"github.com/graph-gophers/graphql-go"
)

type pluginResolver struct {
	id     graphql.ID
	plugin plugin.Plugin
}

func (r *pluginResolver) Id() graphql.ID {
	return r.id
}

func (r *pluginResolver) Name() string {
	return string(r.plugin.Name())
}

func (r *pluginResolver) Description() string {
	return r.plugin.Description()
}
