package gql

import (
	"sort"

	"github.com/adyatlov/xp/data"
	"github.com/graph-gophers/graphql-go"
)

type childrenConnectionResolver struct {
	totalCount  int32
	edges       []*childEdgeResolver
	hasNextPage bool
}

func newChildrenConnectionResolver(
	dId datasetId,
	g data.ObjectGroup,
	first *int32,
	after *graphql.ID) (*childrenConnectionResolver, error) {
	objects := objectPool.Get().(*[]data.Object)
	defer objectPool.Put(objects)
	*objects = (*objects)[:0]
	if err := g.All(objects); err != nil {
		return nil, err
	}
	sort.Slice(*objects, func(i, j int) bool {
		return (*objects)[i].Name() < (*objects)[j].Name()
	})
	totalCount := len(*objects)
	from := 0
	edges := make([]*childEdgeResolver, 0, len(*objects))
	for i, o := range *objects {
		oId := objectId{
			datasetId:      dId,
			ObjectTypeName: o.Type().Name,
			ObjectId:       o.Id(),
		}
		cursor := encodeId(oId)
		edges = append(edges, &childEdgeResolver{
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
	return &childrenConnectionResolver{
		edges:       edges,
		totalCount:  int32(totalCount),
		hasNextPage: hasNextPage,
	}, nil
}

func (r *childrenConnectionResolver) TotalCount() int32 {
	return r.totalCount
}

func (r *childrenConnectionResolver) Edges() *[]*childEdgeResolver {
	return &r.edges
}

func (r *childrenConnectionResolver) PageInfo() pageInfoResolver {
	if len(r.edges) == 0 {
		return pageInfoResolver{}
	}
	return pageInfoResolver{
		startCursor: r.edges[0].cursor,
		endCursor:   r.edges[len(r.edges)-1].cursor,
		hasNextPage: r.hasNextPage,
	}
}

type childEdgeResolver struct {
	cursor graphql.ID
	node   *objectResolver
}

func (r *childEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *childEdgeResolver) Node() *objectResolver {
	return r.node
}
