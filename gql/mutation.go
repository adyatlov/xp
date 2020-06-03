package gql

import (
	"time"

	"github.com/adyatlov/xp/data"
	"github.com/adyatlov/xp/plugin"
)

type Mutation struct {
	datasets        *datasetRegistry
	onDatasetUpdate func()
}

func (m *Mutation) AddDataset(args struct {
	Plugin string
	Url    string
}) (*datasetResolver, error) {
	plugin, err := plugin.GetPlugin(data.PluginName(args.Plugin))
	if err != nil {
		return nil, err
	}
	dataset, err := plugin.Open(args.Url)
	if err != nil {
		return nil, err
	}
	datasetInfo := DatasetInfo{
		Dataset: dataset,
		Plugin:  plugin,
		Url:     args.Url,
		Added:   time.Now(),
	}
	if err := m.datasets.Add(datasetInfo); err != nil {
		return nil, err
	}
	m.onDatasetUpdate()
	return &datasetResolver{datasetInfo}, nil
}

func (m *Mutation) RemoveDataset(args struct {
	Id string
}) (bool, error) {
	if err := m.datasets.Remove(data.DatasetId(args.Id)); err != nil {
		return false, err
	}
	m.onDatasetUpdate()
	return true, nil
}
