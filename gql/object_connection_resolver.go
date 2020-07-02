package gql

import (
	"github.com/adyatlov/xp/data"
	"github.com/graph-gophers/graphql-go"
)

type objectConnectionResolver struct {
	edges []*objectEdgeResolver
}

func newObjectConnectionResolver(
	dId datasetId,
	g data.ObjectGroup,
	first *int32,
	after *graphql.ID) (*objectConnectionResolver, error) {
	if first == nil || *first <= 0 {
		return nil, nil
	}
	objects := objectPool.Get().(*[]data.Object)
	defer objectPool.Put(objects)
	*objects = (*objects)[:0]
	if err := g.All(objects); err != nil {
		return nil, err
	}
	afterIndex := -1
	edges := make([]*objectEdgeResolver, 0, len(*objects))
	if after != nil {
		for i, o := range *objects {
			oId := objectId{
				datasetId:      dId,
				ObjectTypeName: o.Type().Name,
				ObjectId:       o.Id(),
			}
			cursor := encodeId(oId)
			edges = append(edges, &objectEdgeResolver{
				cursor: cursor,
				node:   &objectResolver{objectId: oId, object: o},
			})
			if cursor == *after {
				afterIndex = i
			}
		}
	}
	low := afterIndex + 1
	high := low + int(*first)
	if high > len(*objects) {
		high = len(*objects)
	}
	edges = edges[low:high]
	return &objectConnectionResolver{edges: edges}, nil
}

func (r *objectConnectionResolver) TotalCount() int32 {
	return int32(len(r.edges))
}

func (r *objectConnectionResolver) Edges() *[]*objectEdgeResolver {
	return &r.edges
}

func (r *objectConnectionResolver) PageInfo() pageInfoResolver {
	return pageInfoResolver{}
}

type objectEdgeResolver struct {
	cursor graphql.ID
	node   *objectResolver
}

func (r *objectEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *objectEdgeResolver) Node() *objectResolver {
	return r.node
}
