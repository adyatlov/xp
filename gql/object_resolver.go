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

func (r *objectResolver) Children(args struct{ TypeNames *[]string }) *[]*objectGroupResolver {
	var names []data.ObjectTypeName
	if args.TypeNames == nil || len(*args.TypeNames) == 0 {
		names = r.object.Type().ChildTypeNames()
	} else {
		names = make([]data.ObjectTypeName, 0, len(*args.TypeNames))
		for _, name := range *args.TypeNames {
			names = append(names, data.ObjectTypeName(name))
		}
	}
	objectGroups := make([]*objectGroupResolver, 0, len(names))
	for _, name := range names {
		objectGroup := r.object.Children(name)
		if objectGroup == nil {
			objectGroups = append(objectGroups, nil)
			continue
		}
		objectGroups = append(objectGroups, &objectGroupResolver{g: objectGroup})
	}
	return &objectGroups
}
