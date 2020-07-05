package gql

import (
	"github.com/adyatlov/xp/data"
	"github.com/graph-gophers/graphql-go"
)

type objectResolver struct {
	objectId objectId
	object   data.Object
}

func (r *objectResolver) Id() graphql.ID {
	return encodeId(r.objectId)
}

func (r *objectResolver) Type() *objectTypeResolver {
	return &objectTypeResolver{r.object.Type()}
}

func (r *objectResolver) Name() string {
	return string(r.object.Name())
}

func (r *objectResolver) Properties(
	args struct {
		TypeNames *[]string
		First     *int32
		After     *graphql.ID
	}) (*propertiesConnectionResolver, error) {
	return newPropertiesConnectionResolver(
		r.objectId,
		r.object,
		args.TypeNames,
		args.First,
		args.After,
	)
}

func (r *objectResolver) Children(args struct{ Indexes *[]int32 }) *[]*objectGroupResolver {
	if args.Indexes == nil || len(*args.Indexes) == 0 {
		objectGroups := make([]*objectGroupResolver, 0, len(r.object.Type().ChildTypeNames()))
		for i, name := range r.object.Type().ChildTypeNames() {
			objectGroup := r.object.Children(name)
			if objectGroup == nil {
				objectGroups = append(objectGroups, nil)
				continue
			}
			objectGroups = append(objectGroups, &objectGroupResolver{
				objectId: r.objectId,
				index:    int32(i),
				g:        objectGroup})
		}
		return &objectGroups
	}
	objectGroups := make([]*objectGroupResolver, 0, len(*args.Indexes))
	for _, index := range *args.Indexes {
		index := int(index)
		if index < 0 || index >= len(r.object.Type().ChildTypeNames()) {
			objectGroups = append(objectGroups, nil)
			continue
		}
		t := r.object.Type().ChildTypes[index]
		objectGroup := r.object.Children(t.Name)
		objectGroups = append(objectGroups, &objectGroupResolver{
			objectId: r.objectId,
			index:    int32(index),
			g:        objectGroup})
	}
	return &objectGroups
}
