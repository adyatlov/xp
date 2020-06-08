package gql

import "github.com/graph-gophers/graphql-go"

type node interface {
	Id() graphql.ID
}

type nodeResolver struct {
	node
}

func (r *nodeResolver) ToObject() (resolver *objectResolver, ok bool) {
	resolver, ok = r.node.(*objectResolver)
	return
}

func (r *nodeResolver) ToProperty() (resolver *propertyResolver, ok bool) {
	resolver, ok = r.node.(*propertyResolver)
	return
}

func (r *nodeResolver) ToObjectGroup() (resolver *objectGroupResolver, ok bool) {
	resolver, ok = r.node.(*objectGroupResolver)
	return
}

func (r *nodeResolver) ToDataset() (resolver *datasetResolver, ok bool) {
	resolver, ok = r.node.(*datasetResolver)
	return
}

func (r *nodeResolver) ToPlugin() (resolver *pluginResolver, ok bool) {
	resolver, ok = r.node.(*pluginResolver)
	return
}
