package gql

import (
	"strconv"

	"github.com/graph-gophers/graphql-go"
)

type datasetResolver struct {
	id      datasetId
	dataset DatasetInfo
}

func (r datasetResolver) Id() graphql.ID {
	return encodeId(r.id)
}

func (r datasetResolver) Root() (*objectResolver, error) {
	root, err := r.dataset.Root()
	if err != nil {
		return nil, err
	}
	id := objectId{
		datasetId:      r.id,
		ObjectTypeName: root.Type().Name,
		ObjectId:       root.Id(),
	}
	return &objectResolver{objectId: id, object: root}, nil
}

func (r datasetResolver) Plugin() *pluginResolver {
	return &pluginResolver{plugin: r.dataset.Plugin}
}

func (r datasetResolver) URL() string {
	return r.dataset.Url
}

func (r datasetResolver) Added() string {
	return strconv.FormatInt(r.dataset.Added.Unix(), 10)
}
