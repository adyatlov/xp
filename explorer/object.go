package explorer

import (
	"fmt"
	"sort"

	"github.com/mesosphere/bun/v2/bundle"
)

type ObjectTypeName string
type ObjectId string
type ObjectName string

type Object struct {
	Type     ObjectTypeName `json:"type"`
	Id       ObjectId       `json:"id"`
	Name     ObjectName     `json:"name"`
	Metrics  []*Metric      `json:"metrics"`
	Children []ObjectGroup  `json:"children,omitempty"`
	Errors   []string       `json:"errors,omitempty"`
}

type ObjectGroup struct {
	Type    ObjectTypeName `json:"type"`
	Objects []*Object      `json:"objects"`
}

type ObjectType struct {
	Name              ObjectTypeName                                        `json:"name"`
	DisplayName       string                                                `json:"displayName"`
	PluralDisplayName string                                                `json:"pluralDisplayName"`
	Description       string                                                `json:"description"`
	Metrics           []MetricTypeName                                      `json:"metrics"`
	DefaultMetrics    []MetricTypeName                                      `json:"defaultMetrics"`
	Find              func(*bundle.Bundle, ObjectId, bool) (*Object, error) `json:"-"`
}

func (t ObjectType) New(b *bundle.Bundle, id ObjectId, withChildren bool, metrics ...MetricTypeName) (*Object, error) {
	object, err := t.Find(b, id, withChildren)
	if err != nil {
		return nil, fmt.Errorf("cannot find object: %s", err.Error())
	}
	object.Type = t.Name
	if object.Id == "" {
		object.Id = id
	}
	sortChildren(object.Children)
	if len(metrics) == 0 {
		metrics = t.DefaultMetrics
	}
	for _, requestedMetric := range metrics {
		ok := false
		for _, registeredMetric := range t.Metrics {
			if requestedMetric != registeredMetric {
				continue
			}
			metricType := MustGetMetricType(requestedMetric)
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

func sortChildren(children []ObjectGroup) {
	types := GetObjectTypes()
	sort.Slice(children, func(i, j int) bool {
		return types[children[i].Type].PluralDisplayName < types[children[j].Type].PluralDisplayName
	})
	for _, c := range children {
		sort.Slice(c.Objects, func(i, j int) bool {
			return c.Objects[i].Name < c.Objects[j].Name
		})
	}
}
