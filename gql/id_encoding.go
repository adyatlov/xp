package gql

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/adyatlov/xp/plugin"

	"github.com/adyatlov/xp/data"
	"github.com/graph-gophers/graphql-go"
)

func encodePart(p string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(p))
}

func decodePart(p string) string {
	partBytes, err := base64.RawURLEncoding.DecodeString(p)
	if err != nil {
		panic(err)
	}
	return string(partBytes)
}

func encodeId(id interface{}) graphql.ID {
	switch id := id.(type) {
	case objectId:
		return joinAndEncodeParts("o",
			string(id.PluginName), string(id.DatasetId), string(id.ObjectTypeName), string(id.ObjectId))
	case propertyId:
		return joinAndEncodeParts("p",
			string(id.PluginName), string(id.DatasetId), string(id.ObjectTypeName), string(id.ObjectId), string(id.PropertyName))
	case objectGroupId:
		return joinAndEncodeParts("g",
			string(id.PluginName), string(id.DatasetId), string(id.ObjectTypeName), string(id.ObjectId), string(id.GroupTypeName))
	case datasetId:
		return joinAndEncodeParts("d", string(id.PluginName), string(id.DatasetId))
	case plugin.Name:
		return joinAndEncodeParts("pl", string(id))
	}
	panic(fmt.Sprintf("Cannot encode ID of type %T", id))
}

func decodeId(id graphql.ID) interface{} {
	parts := strings.Split(string(id), ":")
	for i := 1; i < len(parts); i++ {
		parts[i] = decodePart(parts[i])
	}
	switch parts[0] {
	case "o":
		return decodeObjectId(parts)
	case "p":
		return decodePropertyId(parts)
	case "g":
		return decodeObjectGroupId(parts)
	case "d":
		return decodeDatasetId(parts)
	case "pl":
		return plugin.Name(parts[1])
	}
	panic("Cannot decode ID: " + id)
}

func decodeObjectId(parts []string) objectId {
	return objectId{
		datasetId:      decodeDatasetId(parts),
		ObjectTypeName: data.ObjectTypeName(parts[3]),
		ObjectId:       data.ObjectId(parts[4]),
	}
}

func decodePropertyId(parts []string) propertyId {
	return propertyId{
		objectId:     decodeObjectId(parts),
		PropertyName: data.PropertyName(parts[5]),
	}
}

func decodeObjectGroupId(parts []string) objectGroupId {
	return objectGroupId{
		objectId:      decodeObjectId(parts),
		GroupTypeName: data.ObjectTypeName(parts[5]),
	}
}

func decodeDatasetId(parts []string) datasetId {
	return datasetId{
		PluginName: plugin.Name(parts[1]),
		DatasetId:  data.DatasetId(parts[2]),
	}
}

type objectId struct {
	datasetId
	data.ObjectTypeName
	data.ObjectId
}

type propertyId struct {
	objectId
	data.PropertyName
}

type objectGroupId struct {
	objectId
	GroupTypeName data.ObjectTypeName
}

type datasetId struct {
	PluginName plugin.Name
	data.DatasetId
}

func joinAndEncodeParts(parts ...string) graphql.ID {
	for i := 1; i < len(parts); i++ {
		parts[i] = encodePart(parts[i])
	}
	return graphql.ID(strings.Join(parts, ":"))
}
