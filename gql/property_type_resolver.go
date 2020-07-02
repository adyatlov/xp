package gql

import (
	"github.com/adyatlov/xp/data"
)

var propertyValueTypes = []string{
	"BOOL",
	"STRING",
	"INTEGER",
	"REAL",
	"PERCENTAGE",
	"VERSION",
	"TIMESTAMP",
	"TYPE",
	"FILE",
	"OBJECT",
}

type propertyTypeResolver struct {
	t *data.PropertyType
}

func (r *propertyTypeResolver) Name() string {
	return string(r.t.Name)
}

func (r *propertyTypeResolver) ValueType() string {
	return propertyValueTypes[r.t.ValueType]
}

func (r *propertyTypeResolver) Description() string {
	return r.t.Description
}
