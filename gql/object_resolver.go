package gql

import (
	"github.com/adyatlov/xp/data"
	"github.com/graph-gophers/graphql-go"
)

type objectResolver struct {
	id     graphql.ID
	object data.Object
}

func (r *objectResolver) Type() *objectTypeResolver {
	return &objectTypeResolver{r.object.Type()}
}

func (r *objectResolver) Id() graphql.ID {
	return r.id
}

func (r *objectResolver) Name() string {
	return string(r.object.Name())
}

func (r *objectResolver) Children(args struct {
	TypeNames *[]string
}) ([]*childrenGroupResolver, error) {
	var typeNames []data.ObjectTypeName
	if args.TypeNames != nil {
		typeNames = make([]data.ObjectTypeName, 0, len(*args.TypeNames))
		for _, typeName := range *args.TypeNames {
			typeNames = append(typeNames, data.ObjectTypeName(typeName))
		}
	}
	groups, err := r.object.Children(typeNames...)
	if err != nil {
		return nil, err
	}
	resolvers := make([]*childrenGroupResolver, 0, len(groups))
	for _, group := range groups {
		id := encodeId(childrenGroupId{
			objectId:      decodeId(r.id).(objectId),
			GroupTypeName: group.Type().Name,
		})
		resolvers = append(resolvers, &childrenGroupResolver{id: id, group: group})
	}
	return resolvers, nil
}

func (r *objectResolver) Properties(args struct {
	TypeNames *[]string
}) ([]*propertyResolver, error) {
	var typeNames []data.PropertyName
	if args.TypeNames != nil {
		typeNames = make([]data.PropertyName, 0, len(*args.TypeNames))
		for _, typeName := range *args.TypeNames {
			typeNames = append(typeNames, data.PropertyName(typeName))
		}
	}
	properties, err := r.object.Properties(typeNames...)
	if err != nil {
		return nil, err
	}
	resolvers := make([]*propertyResolver, 0, len(properties))
	for _, property := range properties {
		id := encodeId(propertyId{
			objectId:     decodeId(r.id).(objectId),
			PropertyName: property.Type().Name,
		})
		resolvers = append(resolvers, &propertyResolver{id: id, property: property})
	}
	return resolvers, nil
}
