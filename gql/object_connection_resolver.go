package gql

import (
	"github.com/adyatlov/xp/data"
	"github.com/graph-gophers/graphql-go"
)

type objectConnectionResolver struct {
	totalCount  int32
	edges       []*objectEdgeResolver
	hasNextPage bool
}

func newObjectConnectionResolver(
	dId datasetId,
	g data.ObjectGroup,
	first *int32,
	after *graphql.ID) (*objectConnectionResolver, error) {
	objects := objectPool.Get().(*[]data.Object)
	defer objectPool.Put(objects)
	*objects = (*objects)[:0]
	if err := g.All(objects); err != nil {
		return nil, err
	}
	totalCount := len(*objects)
	from := 0
	edges := make([]*objectEdgeResolver, 0, len(*objects))
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
		if after != nil && cursor == *after {
			from = i + 1
		}
	}
	to := len(edges)
	if first != nil {
		to = from + int(*first)
		if to > len(edges) {
			to = len(edges)
		}
	}
	hasNextPage := to < len(edges)
	edges = edges[from:to]
	return &objectConnectionResolver{
		edges:       edges,
		totalCount:  int32(totalCount),
		hasNextPage: hasNextPage,
	}, nil
}

func (r *objectConnectionResolver) TotalCount() int32 {
	return r.totalCount
}

func (r *objectConnectionResolver) Edges() *[]*objectEdgeResolver {
	return &r.edges
}

func (r *objectConnectionResolver) PageInfo() pageInfoResolver {
	if len(r.edges) == 0 {
		return pageInfoResolver{}
	}
	return pageInfoResolver{
		startCursor: r.edges[0].cursor,
		endCursor:   r.edges[len(r.edges)-1].cursor,
		hasNextPage: r.hasNextPage,
	}
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
