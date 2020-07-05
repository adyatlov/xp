package gql

import (
	"github.com/adyatlov/xp/data"
	"github.com/graph-gophers/graphql-go"
)

type objectGroupResolver struct {
	objectId objectId
	index    int32
	g        data.ObjectGroup
}

func (r *objectGroupResolver) Id() graphql.ID {
	return encodeId(objectGroupId{
		objectId:       r.objectId,
		ObjectTypeName: r.g.Type().Name,
	})
}

func (r *objectGroupResolver) Index() int32 {
	return r.index
}

func (r *objectGroupResolver) Type() objectTypeResolver {
	return objectTypeResolver{t: r.g.Type()}
}

func (r *objectGroupResolver) TotalCount() int32 {
	return int32(r.g.TotalCount())
}

func (r *objectGroupResolver) Objects(args struct {
	First *int32
	After *graphql.ID
}) (*objectConnectionResolver, error) {
	return newObjectConnectionResolver(
		r.objectId.datasetId,
		r.g,
		args.First,
		args.After)
}
