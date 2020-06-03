package gql

import (
	"strconv"

	"github.com/graph-gophers/graphql-go"
)

type datasetResolver struct {
	dataset DatasetInfo
}

func (r *datasetResolver) Id() graphql.ID {
	return graphql.ID(r.dataset.Id())
}

func (r *datasetResolver) Root() (*objectResolver, error) {
	root, err := r.dataset.Root()
	if err != nil {
		return nil, err
	}
	return &objectResolver{root}, nil
}

func (r *datasetResolver) Plugin() *pluginResolver {
	return &pluginResolver{r.dataset.Plugin}
}

func (r *datasetResolver) URL() string {
	return r.dataset.Url
}

func (r *datasetResolver) Added() string {
	return strconv.FormatInt(r.dataset.Added.Unix(), 10)
}
