package gql

import (
	"github.com/adyatlov/xp/data"
	"github.com/graph-gophers/graphql-go"
)

type childrenGroupResolver struct {
	group data.ChildrenGroup
	id    graphql.ID
}

func (r *childrenGroupResolver) Id() graphql.ID {
	return r.id
}

func (r *childrenGroupResolver) Type() *objectTypeResolver {
	return &objectTypeResolver{r.group.Type()}
}

func (r *childrenGroupResolver) Objects() []*objectResolver {
	objectResolvers := make([]*objectResolver, 0, len(r.group.Objects()))
	for _, object := range r.group.Objects() {
		id := encodeId(objectId{
			datasetId:      decodeId(r.id).(childrenGroupId).datasetId,
			ObjectTypeName: object.Type().Name,
			ObjectId:       object.Id(),
		})
		objectResolvers = append(objectResolvers, &objectResolver{id: id, object: object})
	}
	return objectResolvers
}

func (r *childrenGroupResolver) Total() int32 {
	return int32(r.group.Total())
}
