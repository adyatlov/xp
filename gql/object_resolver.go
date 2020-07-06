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

func (r *objectResolver) ChildGroups(args struct{ TypeNames *[]string }) *[]*childGroupResolver {
	var names []data.ObjectTypeName
	if args.TypeNames == nil || len(*args.TypeNames) == 0 {
		names = make([]data.ObjectTypeName, 0, len(r.object.Type().ChildTypeNames()))
		for _, name := range r.object.Type().ChildTypeNames() {
			names = append(names, name)
		}
	} else {
		names = make([]data.ObjectTypeName, 0, len(*args.TypeNames))
		for _, name := range *args.TypeNames {
			names = append(names, data.ObjectTypeName(name))
		}
	}
	childGroups := make([]*childGroupResolver, 0, len(names))
	for _, name := range names {
		childGroup := r.object.ChildGroup(name)
		if childGroup == nil {
			childGroups = append(childGroups)
			continue
		}
		childGroups = append(childGroups, &childGroupResolver{
			parentId: r.objectId,
			g:        childGroup})
	}
	return &childGroups
}

func (r *objectResolver) ChildGroup(args struct{ TypeName *string }) *childGroupResolver {
	var typeName data.ObjectTypeName
	if args.TypeName == nil {
		return nil
	}
	typeName = data.ObjectTypeName(*args.TypeName)
	g := r.object.ChildGroup(typeName)
	if g == nil {
		return nil
	}
	return &childGroupResolver{
		parentId: r.objectId,
		g:        g,
	}
}

func (r *objectResolver) FirstAvailableChildGroupTypeName() *string {
	for _, name := range r.object.Type().ChildTypeNames() {
		childGroup := r.object.ChildGroup(name)
		if childGroup == nil {
			continue
		}
		firstAvailableChildGroupTypeName := string(childGroup.Type().Name)
		return &firstAvailableChildGroupTypeName
	}
	return nil
}
