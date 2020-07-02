package gql

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/adyatlov/xp/data"
	"github.com/graph-gophers/graphql-go"
)

type objectId struct {
	datasetId
	data.ObjectTypeName
	data.ObjectId
}

func newObjectId(
	pl data.PluginName,
	d data.DatasetId,
	oT data.ObjectTypeName,
	oId data.ObjectId) objectId {
	dId := newDatasetId(pl, d)
	return objectId{
		datasetId:      dId,
		ObjectTypeName: oT,
		ObjectId:       oId,
	}
}

type propertyId struct {
	objectId
	data.PropertyName
}

func newPropertyId(
	pl data.PluginName,
	d data.DatasetId,
	oT data.ObjectTypeName,
	oId data.ObjectId,
	p data.PropertyName) propertyId {
	objId := newObjectId(pl, d, oT, oId)
	return propertyId{
		objectId:     objId,
		PropertyName: p,
	}
}

type objectGroupId struct {
	objectId
	data.ObjectTypeName
}

func newObjectGroupId(
	pl data.PluginName,
	d data.DatasetId,
	oT data.ObjectTypeName,
	oId data.ObjectId,
	g data.ObjectTypeName) objectGroupId {
	objId := newObjectId(pl, d, oT, oId)
	return objectGroupId{
		objectId:       objId,
		ObjectTypeName: g,
	}
}

type datasetId struct {
	pluginId
	data.DatasetId
}

func newDatasetId(
	pl data.PluginName,
	d data.DatasetId) datasetId {
	plId := newPluginId(pl)
	return datasetId{
		pluginId:  plId,
		DatasetId: d,
	}
}

type pluginId struct {
	PluginName data.PluginName
}

func newPluginId(pl data.PluginName) pluginId {
	return pluginId{pl}
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
			string(id.PluginName), string(id.DatasetId), string(id.ObjectTypeName), string(id.ObjectId), string(id.ObjectTypeName))
	case datasetId:
		return joinAndEncodeParts("d", string(id.PluginName), string(id.DatasetId))
	case pluginId:
		return joinAndEncodeParts("pl", string(id.PluginName))
	}
	panic(fmt.Sprintf("Cannot encode ID of type %T", id))
}

func decodeId(id graphql.ID) (interface{}, error) {
	parts := strings.SplitN(string(id), ":", 2)
	if len(parts) == 1 {
		return nil, fmt.Errorf("wrong ID format: %v", id)
	}
	switch parts[0] {
	case "o":
		return decodeObjectId(id)
	case "p":
		return decodePropertyId(id)
	case "g":
		return decodeObjectGroupId(id)
	case "d":
		return decodeDatasetId(id)
	case "pl":
		return decodePluginId(id)
	}
	return nil, fmt.Errorf("cannot decode ID: %v; unknown prefix: %v", id, parts[0])
}

func decodeObjectId(id graphql.ID) (objectId, error) {
	parts, err := splitAndDecodeParts(id, "o", 5)
	if err != nil {
		return objectId{}, err
	}
	return newObjectId(
		data.PluginName(parts[1]),
		data.DatasetId(parts[2]),
		data.ObjectTypeName(parts[3]),
		data.ObjectId(parts[4])), nil
}

func decodePropertyId(id graphql.ID) (propertyId, error) {
	parts, err := splitAndDecodeParts(id, "p", 6)
	if err != nil {
		return propertyId{}, err
	}
	return newPropertyId(
		data.PluginName(parts[1]),
		data.DatasetId(parts[2]),
		data.ObjectTypeName(parts[3]),
		data.ObjectId(parts[4]),
		data.PropertyName(parts[5])), nil
}

func decodeObjectGroupId(id graphql.ID) (objectGroupId, error) {
	parts, err := splitAndDecodeParts(id, "g", 6)
	if err != nil {
		return objectGroupId{}, err
	}
	return newObjectGroupId(
		data.PluginName(parts[1]),
		data.DatasetId(parts[2]),
		data.ObjectTypeName(parts[3]),
		data.ObjectId(parts[4]),
		data.ObjectTypeName(parts[5])), nil
}

func decodeDatasetId(id graphql.ID) (datasetId, error) {
	parts, err := splitAndDecodeParts(id, "d", 3)
	if err != nil {
		return datasetId{}, err
	}
	return datasetId{
		pluginId:  pluginId{data.PluginName(parts[1])},
		DatasetId: data.DatasetId(parts[2]),
	}, nil
}

func decodePluginId(id graphql.ID) (pluginId, error) {
	parts, err := splitAndDecodeParts(id, "pl", 2)
	if err != nil {
		return pluginId{}, err
	}
	return pluginId{data.PluginName(parts[1])}, nil
}

func joinAndEncodeParts(parts ...string) graphql.ID {
	for i := 1; i < len(parts); i++ {
		parts[i] = encodePart(parts[i])
	}
	return graphql.ID(strings.Join(parts, ":"))
}

func splitAndDecodeParts(id graphql.ID, expectedPrefix string, expectedParts int) ([]string, error) {
	if !strings.HasPrefix(string(id), expectedPrefix) {
		return nil, fmt.Errorf("wrong ID format: %q; expected prefix %q", id, expectedPrefix)
	}
	parts := strings.Split(string(id), ":")
	if len(parts) != expectedParts {
		return nil, fmt.Errorf("wrong ID format: %q; expected %q parts", id, expectedParts)
	}
	for i := 1; i < len(parts); i++ {
		parts[i] = decodePart(parts[i])
	}
	return parts, nil
}

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
