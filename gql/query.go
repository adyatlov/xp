package gql

import (
	"github.com/graph-gophers/graphql-go"

	"github.com/adyatlov/xp/data"
)

type Query struct {
	datasets *datasetRegistry
}

func (q *Query) Object(args struct {
	DatasetId graphql.ID
	Id        graphql.ID
}) (*objectResolver, error) {
	dataset, err := q.datasets.Get(data.DatasetId(args.DatasetId))
	if err != nil {
		return nil, err
	}
	t, id, err := decodeUniqueId(args.Id)
	if err != nil {
		return nil, err
	}
	object, err := dataset.Find(t, id)
	if err != nil {
		return nil, err
	}
	return &objectResolver{object: object}, nil
}

func (q *Query) Datasets(args struct {
	Ids *[]graphql.ID
}) ([]*datasetResolver, error) {
	datasets := q.datasets.GetAll()
	if args.Ids == nil {
		resolvers := make([]*datasetResolver, 0, len(datasets))
		for _, dataset := range datasets {
			resolvers = append(resolvers, &datasetResolver{dataset})
		}
		return resolvers, nil
	}
	resolvers := make([]*datasetResolver, 0, len(*args.Ids))
	for _, id := range *args.Ids {
		dataset, err := q.datasets.Get(data.DatasetId(id))
		if err != nil {
			return nil, err
		}
		resolvers = append(resolvers, &datasetResolver{dataset})
	}
	return resolvers, nil
}

func (q *Query) Plugins(args struct{ Url *string }) ([]*pluginResolver, error) {
	var plugins []data.Plugin
	if args.Url == nil {
		plugins = data.GetPlugins()
	} else {
		var err error
		plugins, err = data.GetCompatiblePlugins(*args.Url)
		if err != nil {
			return nil, err
		}
	}
	resolvers := make([]*pluginResolver, 0, len(plugins))
	for _, plugin := range plugins {
		resolvers = append(resolvers, &pluginResolver{plugin})
	}
	return resolvers, nil
}
