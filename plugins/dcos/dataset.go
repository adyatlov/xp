package dcos

import (
	"fmt"

	"github.com/adyatlov/xp/data"
	"github.com/mesosphere/bun/v2/bundle"
)

type Dataset struct {
	id   data.DatasetId
	b    bundle.Bundle
	root data.Object
}

func (d *Dataset) Id() data.DatasetId {
	return d.id
}

func (d *Dataset) Root() (data.Object, error) {
	var err error
	if d.root == nil {
		d.root, err = newCluster(d.b)
	}
	return d.root, err
}

func (d *Dataset) Find(t data.ObjectTypeName, id data.ObjectId) (data.Object, error) {
	switch t {
	case TCluster.Name:
		return d.Root()
	}
	return nil, fmt.Errorf("no object with type %q and id %q", t, id)
}
