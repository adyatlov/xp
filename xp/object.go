package xp

import (
	"fmt"
	"log"
	"sort"

	"github.com/mesosphere/bun/v2/bundle"
)

type ObjectTypeName string
type ObjectId string
type ObjectName string

type Object interface {
	Type() ObjectTypeName
	Id() ObjectId
	Name() ObjectName
	Metrics(...MetricTypeName) ([]*Metric, error)
	Children() ([]ObjectGroup, error)
}

type ObjectGroup struct {
	Type    ObjectTypeName
	Objects []Object
}

type ObjectType struct {
	Name              ObjectTypeName
	DisplayName       string
	PluralDisplayName string
	Description       string
	Metrics           []MetricTypeName
	DefaultMetrics    []MetricTypeName
	FindObject        func(*bundle.Bundle, ObjectId) (ObjectName, error)
	GetChildren       func(*bundle.Bundle, ObjectId) ([]ObjectGroup, error)
}

func (t ObjectType) New(b *bundle.Bundle, id ObjectId) (Object, error) {
	name, err := t.FindObject(b, id)
	if err != nil {
		return nil, fmt.Errorf("cannot find object: %s", err.Error())
	}
	object := &objectImpl{
		t:    t.Name,
		id:   id,
		name: name,
		children: func() ([]ObjectGroup, error) {
			return t.GetChildren(b, id)
		},
		metrics: func(mm []MetricTypeName) ([]*Metric, error) {
			return t.getMetrics(b, id, mm)
		},
	}
	return object, nil
}

func (t ObjectType) getMetrics(b *bundle.Bundle,
	id ObjectId,
	mm []MetricTypeName) ([]*Metric, error) {
	metrics := make([]*Metric, 0, len(mm))
	for _, requestedMetric := range mm {
		found := false
		for _, registeredMetric := range t.Metrics {
			if requestedMetric != registeredMetric {
				continue
			}
			found = true
			metricType := MustGetMetricType(registeredMetric)
			metric, err := metricType.New(b, id)
			if err != nil {
				log.Printf("cannot create metric: \"%v\": %s\n",
					requestedMetric, err.Error())
				break
			}
			metrics = append(metrics, metric)
			break
		}
		if !found {
			return metrics,
				fmt.Errorf("reguested metric \"%v\" doesn't exist for object type \"%v\"",
					requestedMetric, t.Name)
		}
	}
	return metrics, nil
}

func sortChildren(children []ObjectGroup) {
	types := GetObjectTypes()
	sort.Slice(children, func(i, j int) bool {
		return types[children[i].Type].PluralDisplayName < types[children[j].Type].PluralDisplayName
	})
	for _, c := range children {
		sort.Slice(c.Objects, func(i, j int) bool {
			return c.Objects[i].Name() < c.Objects[j].Name()
		})
	}
}

type objectImpl struct {
	t        ObjectTypeName
	id       ObjectId
	name     ObjectName
	children func() ([]ObjectGroup, error)
	metrics  func([]MetricTypeName) ([]*Metric, error)
}

func (o *objectImpl) Type() ObjectTypeName {
	return o.t
}

func (o *objectImpl) Id() ObjectId {
	return o.id
}

func (o *objectImpl) Name() ObjectName {
	return o.name
}

func (o *objectImpl) Metrics(mm ...MetricTypeName) ([]*Metric, error) {
	return o.metrics(mm)
}

func (o *objectImpl) Children() ([]ObjectGroup, error) {
	return o.children()
}
