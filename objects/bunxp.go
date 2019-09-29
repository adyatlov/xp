package objects

import (
	"fmt"

	"github.com/mesosphere/bun/v2/bundle"
)

type ObjectExplorer struct {
	Bundle *bundle.Bundle
}

func (b *ObjectExplorer) Object(n ObjectTypeName, id ObjectId, metrics ...MetricName) (*Object, error) {
	t, err := GetObjectType(n)
	if err != nil {
		return nil, fmt.Errorf("cannot create object: %s", err.Error())
	}
	return t.New(b.Bundle, id, metrics...)
}
