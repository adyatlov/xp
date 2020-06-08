package gql

import (
	"strings"

	"github.com/adyatlov/xp/data"
)

type propertyTypeResolver struct {
	t *data.PropertyType
}

func (r *propertyTypeResolver) Name() string {
	return string(r.t.Name)
}

func (r *propertyTypeResolver) ValueType() string {
	return strings.ToUpper(string(r.t.Type))
}

func (r *propertyTypeResolver) Description() string {
	return r.t.Description
}
