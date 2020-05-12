package gql

import (
	"github.com/adyatlov/xp/data"
	"github.com/graph-gophers/graphql-go"
)

type datasetResolver struct {
	dataset data.Dataset
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
