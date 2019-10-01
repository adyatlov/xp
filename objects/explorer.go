package objects

import (
	"fmt"

	"github.com/mesosphere/bun/v2/bundle"
)

type Explorer struct {
	Bundle *bundle.Bundle
}

func (b *Explorer) Object(n ObjectTypeName, id ObjectId, metrics ...MetricTypeName) (*Object, error) {
	t, err := GetObjectType(n)
	if err != nil {
		return nil, fmt.Errorf("cannot create object: %s", err.Error())
	}
	return t.New(b.Bundle, id, metrics...)
}
