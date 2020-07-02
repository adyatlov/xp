package gql

import (
	"time"

	"github.com/adyatlov/xp/data"

	"github.com/graph-gophers/graphql-go"
)

type Mutation struct {
	datasets         *datasetRegistry
	onDatasetAdded   func(r *datasetResolver)
	onDatasetRemoved func(id graphql.ID)
}

func (m *Mutation) AddDataset(args struct {
	PluginName string
	Url        string
}) (*datasetResolver, error) {
	plugin, err := data.GetPlugin(data.PluginName(args.PluginName))
	if err != nil {
		return nil, err
	}
	dataset, err := plugin.Open(args.Url)
	if err != nil {
		return nil, err
	}
	datasetInfo := DatasetInfo{
		Plugin:  plugin,
		Dataset: dataset,
		Url:     args.Url,
		Added:   time.Now(),
	}
	if err := m.datasets.Add(datasetInfo); err != nil {
		return nil, err
	}
	id := newDatasetId(plugin.Name, dataset.Id())
	r := &datasetResolver{id: id, dataset: datasetInfo}
	m.onDatasetAdded(r)
	return r, nil
}

func (m *Mutation) RemoveDataset(args struct{ Id graphql.ID }) (bool, error) {
	id, err := decodeId(args.Id)
	if err != nil {
		return false, err
	}
	if err := m.datasets.Remove(id.(datasetId)); err != nil {
		return false, err
	}
	m.onDatasetRemoved(args.Id)
	return true, nil
}
