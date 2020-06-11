package gql

import (
	"fmt"

	"github.com/adyatlov/xp/plugin"
	"github.com/graph-gophers/graphql-go"
)

type Query struct {
	datasets *datasetRegistry
}

func (q *Query) Node(args struct {
	Id *graphql.ID
}) (*nodeResolver, error) {
	if args.Id == nil {
		return nil, nil
	}
	id := decodeId(*args.Id)
	switch id := id.(type) {
	case objectId:
		dataset, err := q.datasets.Get(id.datasetId)
		if err != nil {
			return nil, err
		}
		object, err := dataset.Find(id.ObjectTypeName, id.ObjectId)
		if err != nil {
			return nil, err
		}
		return &nodeResolver{&objectResolver{id: *args.Id, object: object}}, nil
	case propertyId:
		dataset, err := q.datasets.Get(id.datasetId)
		if err != nil {
			return nil, err
		}
		object, err := dataset.Find(id.ObjectTypeName, id.ObjectId)
		if err != nil {
			return nil, err
		}
		properties, err := object.Properties(id.PropertyName)
		if err != nil {
			return nil, err
		}
		return &nodeResolver{&propertyResolver{id: *args.Id, property: properties[0]}}, nil
	case childrenGroupId:
		dataset, err := q.datasets.Get(id.datasetId)
		if err != nil {
			return nil, err
		}
		object, err := dataset.Find(id.ObjectTypeName, id.ObjectId)
		if err != nil {
			return nil, err
		}
		groups, err := object.Children(id.GroupTypeName)
		if err != nil {
			return nil, err
		}
		return &nodeResolver{&childrenGroupResolver{id: *args.Id, group: groups[0]}}, nil
	case datasetId:
		dataset, err := q.datasets.Get(id)
		if err != nil {
			return nil, err
		}
		return &nodeResolver{&datasetResolver{id: *args.Id, dataset: dataset}}, nil
	case plugin.Name:
		p, err := plugin.GetPlugin(id)
		if err != nil {
			return nil, err
		}
		return &nodeResolver{&pluginResolver{id: *args.Id, plugin: p}}, nil
	}
	panic(fmt.Sprintf("Unknown ID type: %T", id))
}

func (q *Query) AllDatasets() []*datasetResolver {
	datasets := q.datasets.GetAll()
	resolvers := make([]*datasetResolver, 0, len(datasets))
	for _, dataset := range datasets {
		id := encodeId(datasetId{
			PluginName: dataset.Plugin.Name(),
			DatasetId:  dataset.Id(),
		})
		resolvers = append(resolvers, &datasetResolver{id: id, dataset: dataset})
	}
	return resolvers
}

func (q *Query) AllPlugins() []*pluginResolver {
	plugins := plugin.GetPlugins()
	resolvers := make([]*pluginResolver, 0, len(plugins))
	for _, p := range plugins {
		id := encodeId(p.Name())
		resolvers = append(resolvers, &pluginResolver{id: id, plugin: p})
	}
	return resolvers
}

func (q *Query) CompatiblePlugins(args struct{ Url *string }) ([]*pluginResolver, error) {
	if args.Url == nil {
		return []*pluginResolver{}, nil
	}
	plugins, err := plugin.GetCompatiblePlugins(*args.Url)
	if err != nil {
		return nil, err
	}
	resolvers := make([]*pluginResolver, 0, len(plugins))
	for _, p := range plugins {
		id := encodeId(p.Name())
		resolvers = append(resolvers, &pluginResolver{id: id, plugin: p})
	}
	return resolvers, nil
}
