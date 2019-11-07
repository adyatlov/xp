package server

import (
	"fmt"
	"strings"

	"github.com/adyatlov/bunxp/explorer"
	"github.com/graph-gophers/graphql-go"
)

type resolver struct {
	explorer *explorer.Explorer
}

func (r *resolver) Roots() (*[]*objectResolver, error) {
	t := explorer.ObjectTypeName("cluster")
	object, err := r.explorer.Object(t, "")
	return &[]*objectResolver{{object: object}}, err
}

func (r *resolver) Object(args struct {
	ObjectId string
	Type     string
}) (*objectResolver, error) {
	t := explorer.ObjectTypeName(args.Type)
	object, err := r.explorer.Object(t, explorer.ObjectId(args.ObjectId))
	return &objectResolver{object: object}, err
}

func (r *resolver) ObjectTypes() []*objectTypeResolver {
	objTypeMap := explorer.GetObjectTypes()
	objTypes := make([]*objectTypeResolver, 0, len(objTypeMap))
	for _, t := range objTypeMap {
		objTypes = append(objTypes, &objectTypeResolver{t: t})
	}
	return objTypes
}

func (r *resolver) MetricTypes() *[]*metricTypeResolver {
	mTypeMap := explorer.GetMetricTypes()
	mTypes := make([]*metricTypeResolver, 0, len(mTypeMap))
	for _, t := range mTypeMap {
		mTypes = append(mTypes, &metricTypeResolver{t: t})
	}
	return &mTypes
}

type objectResolver struct {
	object *explorer.Object
}

func (r *objectResolver) ID() graphql.ID {
	return graphql.ID(string(r.object.Type) + ":::" + string(r.object.Id))
}

func (r *objectResolver) ObjectId() string {
	return string(r.object.Id)
}

func (r *objectResolver) Type() string {
	return string(r.object.Type)
}

func (r *objectResolver) Name() string {
	return string(r.object.Name)
}

func (r *objectResolver) Metrics() *[]*metricResolver {
	metricResolvers := make([]*metricResolver, 0, len(r.object.Metrics))
	for _, metric := range r.object.Metrics {
		metricResolvers = append(metricResolvers, &metricResolver{metric: metric})
	}
	return &metricResolvers
}

func (r *objectResolver) Children() *[]*objectGroupResolver {
	objectGroupResolvers := make([]*objectGroupResolver, 0, len(r.object.Children))
	for _, group := range r.object.Children {
		objectGroupResolvers = append(objectGroupResolvers, &objectGroupResolver{group: &group})
	}
	return &objectGroupResolvers
}

type objectGroupResolver struct {
	group *explorer.ObjectGroup
}

func (r *objectGroupResolver) Type() string {
	return string(r.group.Type)
}

func (r *objectGroupResolver) Objects() *[]*objectResolver {
	objectResolvers := make([]*objectResolver, 0, len(r.group.Objects))
	for _, object := range r.group.Objects {
		objectResolvers = append(objectResolvers, &objectResolver{object: object})
	}
	return &objectResolvers
}

type metricResolver struct {
	metric *explorer.Metric
}

func (r *metricResolver) Type() string {
	return string(r.metric.Type)
}

func (r *metricResolver) Value() string {
	return fmt.Sprintf("%v", r.metric.Value)
}

type objectTypeResolver struct {
	t explorer.ObjectType
}

func (r *objectTypeResolver) Name() string {
	return string(r.t.Name)
}

func (r *objectTypeResolver) DisplayName() string {
	return r.t.DisplayName
}

func (r *objectTypeResolver) PluralDisplayName() string {
	return r.t.PluralDisplayName
}

func (r *objectTypeResolver) Description() string {
	return r.t.Description
}

func (r *objectTypeResolver) Metrics() *[]string {
	metrics := make([]string, 0, len(r.t.Metrics))
	for _, m := range r.t.Metrics {
		metrics = append(metrics, string(m))
	}
	return &metrics
}

func (r *objectTypeResolver) DefaultMetrics() *[]string {
	metrics := make([]string, 0, len(r.t.DefaultMetrics))
	for _, m := range r.t.DefaultMetrics {
		metrics = append(metrics, string(m))
	}
	return &metrics
}

type metricTypeResolver struct {
	t explorer.MetricType
}

func (r *metricTypeResolver) Name() string {
	return string(r.t.Name)
}

func (r *metricTypeResolver) ObjectTypeName() string {
	return string(r.t.ObjectTypeName)
}

func (r *metricTypeResolver) ValueType() string {
	return strings.ToUpper(string(r.t.ValueType))
}

func (r *metricTypeResolver) MetricName() string {
	return string(r.t.MetricName)
}

func (r *metricTypeResolver) Description() string {
	return r.t.Description
}
