package gql

import (
	"github.com/adyatlov/xp/data"
	"github.com/graph-gophers/graphql-go"
)

type propertiesConnectionResolver struct {
	objectId    objectId
	object      data.Object
	names       []data.PropertyName
	hasNextPage bool
}

func newPropertiesConnectionResolver(
	oId objectId,
	object data.Object,
	typeNames *[]string,
	first *int32,
	after *graphql.ID) (*propertiesConnectionResolver, error) {
	if first == nil || *first <= 0 {
		return nil, nil
	}
	var names []data.PropertyName
	if typeNames == nil || len(*typeNames) == 0 {
		names = object.Type().PropertyNames() // Return all properties
	} else {
		names = make([]data.PropertyName, 0, len(*typeNames))
		for _, name := range *typeNames {
			names = append(names, data.PropertyName(name))
		}
	}
	afterIndex := -1
	// Find if object has a property to which "after" is pointing.
	if after != nil {
		afterId, err := decodePropertyId(*after)
		if err != nil {
			return nil, err
		}
		if afterId.objectId != oId {
			// afterId is valid but it points to a property of a another object
			goto Paging
		}
		for i, name := range names {
			if afterId.PropertyName == name {
				afterIndex = i
				goto Paging
			}
		}
	}
Paging:
	low := afterIndex + 1
	high := low + int(*first)
	if high > len(names) {
		high = len(names)
	}
	hasNextPage := high < len(names)
	names = names[low:high]
	if len(names) == 0 {
		return nil, nil
	}
	return &propertiesConnectionResolver{
		objectId:    oId,
		object:      object,
		names:       names,
		hasNextPage: hasNextPage,
	}, nil
}

func (r *propertiesConnectionResolver) TotalCount() int32 {
	return int32(len(r.names))
}

func (r *propertiesConnectionResolver) Edges() (*[]*propertiesEdgeResolver, error) {
	properties := propertyPool.Get().(*[]interface{})
	defer propertyPool.Put(properties)
	*properties = (*properties)[:0]
	err := r.object.Properties(properties, r.names...)
	if err != nil {
		return nil, err
	}
	resolvers := make([]*propertiesEdgeResolver, 0, len(r.names))
	for i, v := range *properties {
		resolvers = append(resolvers, &propertiesEdgeResolver{
			dId:    r.objectId.datasetId,
			cursor: encodeId(propertyId{r.objectId, r.names[i]}),
			t:      r.object.Type().PropertyType(r.names[i]),
			v:      v,
		})
	}
	return &resolvers, nil
}

func (r *propertiesConnectionResolver) PageInfo() pageInfoResolver {
	startCursor := encodeId(propertyId{r.objectId, r.names[0]})
	endCursor := encodeId(propertyId{r.objectId, r.names[len(r.names)-1]})
	return pageInfoResolver{
		startCursor: startCursor,
		endCursor:   endCursor,
		hasNextPage: r.hasNextPage,
	}
}

type propertiesEdgeResolver struct {
	dId    datasetId
	cursor graphql.ID
	t      *data.PropertyType
	v      interface{}
}

func (r *propertiesEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *propertiesEdgeResolver) Node() *propertyResolver {
	return &propertyResolver{dId: r.dId, id: r.cursor, t: r.t, v: r.v}
}
