package objects

import (
	"fmt"

	"github.com/mesosphere/bun/v2/bundle"
)

type ObjectTypeName string
type ObjectId string
type ObjectName string

type Object struct {
	Name     ObjectName
	Type     ObjectTypeName
	Id       ObjectId
	Metrics  []*Metric
	Children map[ObjectTypeName][]ObjectId
	Errors   []string
}

type ObjectType struct {
	Name           ObjectTypeName
	Description    string
	Metrics        []*MetricType
	DefaultMetrics []MetricTypeName
	Find           func(*bundle.Bundle, ObjectId) (*Object, error)
	Children       func(*bundle.Bundle, ObjectId) (map[ObjectTypeName][]ObjectId, error)
}

func (t ObjectType) New(b *bundle.Bundle, id ObjectId, metrics ...MetricTypeName) (*Object, error) {
	object, err := t.Find(b, id)
	if err != nil {
		return nil, fmt.Errorf("cannot find object: %s", err.Error())
	}
	object.Type = t.Name
	if object.Id == "" {
		object.Id = id
	}
	if t.Children != nil {
		if object.Children, err = t.Children(b, id); err != nil {
			object.Errors = append(object.Errors, fmt.Sprintf("cannot find children: %s", err.Error()))
		}
	}
	if len(metrics) == 0 {
		metrics = t.DefaultMetrics
	}
	for _, requestedMetric := range metrics {
		ok := false
		for _, metricType := range t.Metrics {
			if requestedMetric != metricType.Name {
				continue
			}
			metric, err := metricType.New(b, object.Id)
			if err != nil {
				object.Errors = append(object.Errors,
					fmt.Sprintf("cannot create metric: \"%v\": %s",
						requestedMetric, err.Error()))
			}
			object.Metrics = append(object.Metrics, metric)
			ok = true
			break
		}
		if !ok {
			object.Errors = append(object.Errors,
				fmt.Sprintf("reguested metric \"%v\" doesn't exist", requestedMetric))
		}
	}
	return object, nil
}
