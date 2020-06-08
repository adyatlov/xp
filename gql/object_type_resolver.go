package gql

import "github.com/adyatlov/xp/data"

type objectTypeResolver struct {
	t *data.ObjectType
}

func (r *objectTypeResolver) Name() string {
	return string(r.t.Name)
}

func (r *objectTypeResolver) PluralName() string {
	return r.t.PluralName
}

func (r *objectTypeResolver) Description() string {
	return r.t.Description
}

func (r *objectTypeResolver) Properties() []*propertyTypeResolver {
	resolvers := make([]*propertyTypeResolver, 0, len(r.t.Properties))
	for _, property := range r.t.Properties {
		resolvers = append(resolvers, &propertyTypeResolver{property})
	}
	return resolvers
}
