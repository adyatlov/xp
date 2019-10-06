package objects

import (
	"fmt"

	"github.com/mesosphere/bun/v2/bundle"
)

type Explorer struct {
	bundle *bundle.Bundle
}

func NewExplorer(b *bundle.Bundle) *Explorer {
	return &Explorer{bundle: b}
}

func (b *Explorer) Object(n ObjectTypeName, id ObjectId, metrics ...MetricTypeName) (*Object, error) {
	t, err := GetObjectType(n)
	if err != nil {
		return nil, fmt.Errorf("cannot create object: %s", err.Error())
	}
	return t.New(b.bundle, id, true, metrics...)
}
