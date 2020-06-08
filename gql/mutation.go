package gql

import (
	"time"

	"github.com/graph-gophers/graphql-go"

	"github.com/adyatlov/xp/plugin"
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
	plugin, err := plugin.GetPlugin(plugin.Name(args.PluginName))
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
	id := encodeId(datasetId{PluginName: plugin.Name(), DatasetId: dataset.Id()})
	r := &datasetResolver{id: id, dataset: datasetInfo}
	m.onDatasetAdded(r)
	return r, nil
}

func (m *Mutation) RemoveDataset(args struct{ Id graphql.ID }) (bool, error) {
	id := decodeId(args.Id).(datasetId)
	if err := m.datasets.Remove(id); err != nil {
		return false, err
	}
	m.onDatasetRemoved(args.Id)
	return true, nil
}
