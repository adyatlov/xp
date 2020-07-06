package gql

import (
	"github.com/adyatlov/xp/data"
	"github.com/graph-gophers/graphql-go"
)

type childGroupResolver struct {
	parentId objectId
	g        data.ObjectGroup
}

func (r *childGroupResolver) Id() graphql.ID {
	return encodeId(objectGroupId{
		objectId:       r.parentId,
		ObjectTypeName: r.g.Type().Name,
	})
}

func (r *childGroupResolver) Type() objectTypeResolver {
	return objectTypeResolver{t: r.g.Type()}
}

func (r *childGroupResolver) TotalCount() int32 {
	return int32(r.g.TotalCount())
}

func (r *childGroupResolver) Children(args struct {
	First *int32
	After *graphql.ID
}) (*childrenConnectionResolver, error) {
	return newChildrenConnectionResolver(
		r.parentId.datasetId,
		r.g,
		args.First,
		args.After)
}
