package gql

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/adyatlov/xp/data"
	"github.com/graph-gophers/graphql-go"
)

func encodeUniqueId(t data.ObjectTypeName, id data.ObjectId) graphql.ID {
	typePart := base64.RawURLEncoding.EncodeToString([]byte(t))
	idPart := base64.RawURLEncoding.EncodeToString([]byte(id))
	return graphql.ID(typePart + ":" + idPart)
}

func decodeUniqueId(id graphql.ID) (data.ObjectTypeName, data.ObjectId, error) {
	typeAndId := strings.Split(string(id), ":")
	if len(typeAndId) != 2 {
		return "", "", fmt.Errorf("wrong ID format: %v", typeAndId)
	}
	typeStr, err := base64.RawURLEncoding.DecodeString(typeAndId[0])
	idStr, err := base64.RawURLEncoding.DecodeString(typeAndId[1])
	if err != nil {
		return "", "", err
	}
	return data.ObjectTypeName(typeStr), data.ObjectId(idStr), nil
}

type objectResolver struct {
	object data.Object
}

func (r *objectResolver) Type() *objectTypeResolver {
	return &objectTypeResolver{r.object.Type()}
}

func (r *objectResolver) Id() graphql.ID {
	return encodeUniqueId(r.object.Type().Name, r.object.Id())
}

func (r *objectResolver) Name() string {
	return string(r.object.Name())
}

func (r *objectResolver) Children(args struct {
	TypeNames *[]string
}) ([]*objectGroupResolver, error) {
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
	resolvers := make([]*objectGroupResolver, 0, len(groups))
	for _, group := range groups {
		resolvers = append(resolvers, &objectGroupResolver{group})
	}
	return resolvers, nil
}

func (r *objectResolver) Properties(args struct {
	TypeNames *[]string
}) ([]*propertyResolver, error) {
	var typeNames []data.PropertyTypeName
	if args.TypeNames != nil {
		typeNames = make([]data.PropertyTypeName, 0, len(*args.TypeNames))
		for _, typeName := range *args.TypeNames {
			typeNames = append(typeNames, data.PropertyTypeName(typeName))
		}
	}
	properties, err := r.object.Properties(typeNames...)
	if err != nil {
		return nil, err
	}
	resolvers := make([]*propertyResolver, 0, len(properties))
	for _, property := range properties {
		resolvers = append(resolvers, &propertyResolver{property})
	}
	return resolvers, nil
}
