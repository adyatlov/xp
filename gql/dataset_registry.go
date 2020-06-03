package gql

import (
	"fmt"
	"sync"
	"time"

	"github.com/adyatlov/xp/data"
	"github.com/adyatlov/xp/plugin"
)

type DatasetInfo struct {
	data.Dataset
	Plugin plugin.Plugin
	Url    string
	Added  time.Time
}

type datasetRegistry struct {
	datasets   map[data.DatasetId]DatasetInfo
	datasetsMu *sync.RWMutex
}

func NewDatasetRegistry() *datasetRegistry {
	registry := &datasetRegistry{}
	registry.datasets = make(map[data.DatasetId]DatasetInfo)
	registry.datasetsMu = &sync.RWMutex{}
	return registry
}

func (r *datasetRegistry) Add(dataset DatasetInfo) error {
	r.datasetsMu.Lock()
	defer r.datasetsMu.Unlock()
	if _, ok := r.datasets[dataset.Id()]; ok {
		return fmt.Errorf("dataset %v already opened", dataset.Id())
	}
	r.datasets[dataset.Id()] = dataset
	return nil
}

func (r *datasetRegistry) Get(id data.DatasetId) (DatasetInfo, error) {
	r.datasetsMu.RLock()
	defer r.datasetsMu.RUnlock()
	if dataset, ok := r.datasets[id]; ok {
		return dataset, nil
	}
	return DatasetInfo{}, fmt.Errorf("dataset \"%v\" not found", id)
}

func (r *datasetRegistry) GetAll() []DatasetInfo {
	r.datasetsMu.RLock()
	defer r.datasetsMu.RUnlock()
	res := make([]DatasetInfo, 0, len(r.datasets))
	for _, dataset := range r.datasets {
		res = append(res, dataset)
	}
	return res
}

func (r *datasetRegistry) Remove(id data.DatasetId) error {
	r.datasetsMu.Lock()
	defer r.datasetsMu.Unlock()
	_, ok := r.datasets[id]
	if !ok {
		return fmt.Errorf("dataset %v doesn't exist", id)
	}
	delete(r.datasets, id)
	return nil
}
