package gql

import (
	"fmt"
	"sort"

	"github.com/adyatlov/xp/data"

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
	id, err := decodeId(*args.Id)
	if err != nil {
		return nil, err
	}
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
		return &nodeResolver{&objectResolver{objectId: id, object: object}}, nil
	case propertyId:
		dataset, err := q.datasets.Get(id.datasetId)
		if err != nil {
			return nil, err
		}
		object, err := dataset.Find(id.ObjectTypeName, id.ObjectId)
		if err != nil {
			return nil, err
		}
		properties := propertyPool.Get().(*[]interface{})
		defer propertyPool.Put(properties)
		*properties = (*properties)[:0]
		err = object.Properties(properties, id.PropertyName)
		if err != nil {
			return nil, err
		}
		t := object.Type().PropertyType(id.PropertyName)
		v := (*properties)[0]
		r := &propertyResolver{
			dId: id.datasetId,
			t:   t,
			v:   v,
			id:  *args.Id,
		}
		return &nodeResolver{r}, nil
	case objectGroupId:
		dataset, err := q.datasets.Get(id.datasetId)
		if err != nil {
			return nil, err
		}
		object, err := dataset.Find(id.ObjectTypeName, id.ObjectId)
		if err != nil {
			return nil, err
		}
		g := object.ChildGroup(id.ObjectTypeName)
		if g == nil {
			return nil, nil
		}
		r := &childGroupResolver{
			parentId: id.objectId,
			g:        g,
		}
		return &nodeResolver{r}, nil
	case datasetId:
		dataset, err := q.datasets.Get(id)
		if err != nil {
			return nil, err
		}
		r := &datasetResolver{
			id:      id,
			dataset: dataset,
		}
		return &nodeResolver{r}, nil
	case pluginId:
		p, err := data.GetPlugin(id.PluginName)
		if err != nil {
			return nil, err
		}
		return &nodeResolver{&pluginResolver{plugin: p}}, nil
	}
	panic(fmt.Sprintf("Unknown ID type: %T", id))
}

func (q *Query) Datasets() *[]datasetResolver {
	datasets := q.datasets.GetAll()
	if len(datasets) == 0 {
		var d []datasetResolver
		return &d
	}
	resolvers := make([]datasetResolver, 0, len(datasets))
	for _, dataset := range datasets {
		id := newDatasetId(dataset.Plugin.Name, dataset.Id())
		resolvers = append(resolvers, datasetResolver{
			id:      id,
			dataset: dataset,
		})
	}
	sort.Slice(resolvers, func(i, j int) bool {
		return resolvers[i].dataset.Added.Before(
			resolvers[j].dataset.Added)
	})
	return &resolvers
}

func (q *Query) Plugins() *[]pluginResolver {
	plugins := data.GetPlugins()
	resolvers := make([]pluginResolver, 0, len(plugins))
	for _, p := range plugins {
		resolvers = append(resolvers, pluginResolver{plugin: p})
	}
	sort.Slice(resolvers, func(i, j int) bool {
		return resolvers[i].plugin.Name < resolvers[j].plugin.Name
	})
	return &resolvers
}

func (q *Query) CompatiblePlugins(args struct{ Url *string }) (*[]pluginResolver, error) {
	if args.Url == nil {
		return nil, nil
	}
	plugins, err := data.GetCompatiblePlugins(*args.Url)
	if err != nil {
		return nil, err
	}
	resolvers := make([]pluginResolver, 0, len(plugins))
	for _, p := range plugins {
		resolvers = append(resolvers, pluginResolver{plugin: p})
	}
	return &resolvers, nil
}
