package objects

import (
	"fmt"

	"github.com/mesosphere/bun/v2/bundle"
)

type ObjectTypeName string
type ObjectId string
type ObjectName string

type Object struct {
	Type     ObjectTypeName
	Id       ObjectId
	Name     ObjectName
	Metrics  []*Metric
	Children []*Object
	Parents  []*Object
	Errors   []string
}

type ObjectType struct {
	Name           ObjectTypeName
	Description    string
	Metrics        []Metric
	DefaultMetrics []MetricName
	Find           func(*bundle.Bundle, ObjectId) (*Object, error)
	Parents        func(*bundle.Bundle, ObjectId) ([]*Object, error)
	Children       func(*bundle.Bundle, ObjectId) ([]*Object, error)
}

func (t ObjectType) New(b *bundle.Bundle, id ObjectId, metrics ...MetricName) (*Object, error) {
	object, err := t.Find(b, id)
	if err != nil {
		return nil, fmt.Errorf("cannot create object: %s", err.Error())
	}
	object.Type = t.Name
	if object.Id == "" && id != "" {
		object.Id = id
	}
	if t.Parents != nil {
		if object.Parents, err = t.Parents(b, id); err != nil {
			object.Errors = append(object.Errors, fmt.Sprintf("cannot find parents: %s", err.Error()))
		}
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
		for _, metric := range t.Metrics {
			if requestedMetric != metric.Name {
				continue
			}
			metric.Value, err = metric.Evaluate(b, object)
			if err != nil {
				object.Errors = append(object.Errors,
					fmt.Sprintf("cannot evaluate metric %v: %s",
						requestedMetric, err.Error()))
			}
			object.Metrics = append(object.Metrics, &metric)
			ok = true
			break
		}
		if !ok {
			object.Errors = append(object.Errors,
				fmt.Sprintf("reguested metric %v doesn't exist", requestedMetric))
		}
	}
	return object, nil
}
