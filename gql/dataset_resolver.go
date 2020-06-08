package gql

import (
	"strconv"

	"github.com/graph-gophers/graphql-go"
)

type datasetResolver struct {
	id      graphql.ID
	dataset DatasetInfo
}

func (r *datasetResolver) Id() graphql.ID {
	return r.id
}

func (r *datasetResolver) Root() (*objectResolver, error) {
	root, err := r.dataset.Root()
	if err != nil {
		return nil, err
	}
	id := encodeId(objectId{
		datasetId:      datasetId{PluginName: r.dataset.Plugin.Name(), DatasetId: r.dataset.Id()},
		ObjectTypeName: root.Type().Name,
		ObjectId:       root.Id(),
	})
	return &objectResolver{id: id, object: root}, nil
}

func (r *datasetResolver) Plugin() *pluginResolver {
	id := encodeId(r.dataset.Plugin.Name())
	return &pluginResolver{id: id, plugin: r.dataset.Plugin}
}

func (r *datasetResolver) URL() string {
	return r.dataset.Url
}

func (r *datasetResolver) Added() string {
	return strconv.FormatInt(r.dataset.Added.Unix(), 10)
}
