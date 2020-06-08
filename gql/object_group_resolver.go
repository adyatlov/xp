package gql

import (
	"github.com/adyatlov/xp/data"
	"github.com/graph-gophers/graphql-go"
)

type objectGroupResolver struct {
	group data.ObjectGroup
	id    graphql.ID
}

func (r *objectGroupResolver) Id() graphql.ID {
	return r.id
}

func (r *objectGroupResolver) Type() *objectTypeResolver {
	return &objectTypeResolver{r.group.Type()}
}

func (r *objectGroupResolver) Objects() []*objectResolver {
	objectResolvers := make([]*objectResolver, 0, len(r.group.Objects()))
	for _, object := range r.group.Objects() {
		objectResolvers = append(objectResolvers, &objectResolver{object: object})
	}
	return objectResolvers
}

func (r *objectGroupResolver) Total() int32 {
	return int32(r.group.Total())
}
