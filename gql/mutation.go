package gql

import "github.com/adyatlov/xp/data"

type Mutation struct {
	datasets *datasetRegistry
}

func (m *Mutation) AddDataset(args struct {
	Plugin string
	Url    string
}) (*datasetResolver, error) {
	plugin, err := data.GetPlugin(data.PluginName(args.Plugin))
	if err != nil {
		return nil, err
	}
	dataset, err := plugin.Open(args.Url)
	if err != nil {
		return nil, err
	}
	if err := m.datasets.Add(dataset); err != nil {
		return nil, err
	}
	return &datasetResolver{dataset}, nil
}

func (m *Mutation) RemoveDataset(args struct {
	Id string
}) (bool, error) {
	if err := m.datasets.Remove(data.DatasetId(args.Id)); err != nil {
		return false, err
	}
	return true, nil
}
