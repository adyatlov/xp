package explorer

import (
	"fmt"

	"github.com/mesosphere/bun/v2/bundle"
)

type MetricValueType string
type MetricTypeName string
type MetricName string

const (
	MTInteger    MetricValueType = "integer"
	MTReal                       = "real"
	MTPercentage                 = "percentage"
	MTVersion                    = "version"
	MTTimestamp                  = "timestamp"
	MTType                       = "type"
)

type Metric struct {
	Type  MetricTypeName `json:"type"`
	Value interface{}    `json:"value"`
}

type MetricType struct {
	Name           MetricTypeName                                      `json:"name"`
	ObjectTypeName ObjectTypeName                                      `json:"objectTypeName"`
	ValueType      MetricValueType                                     `json:"valueType"`
	MetricName     MetricName                                          `json:"metricName"`
	Description    string                                              `json:"description"`
	Evaluate       func(*bundle.Bundle, ObjectId) (interface{}, error) `json:"-"`
}

func (t MetricType) New(b *bundle.Bundle, id ObjectId) (*Metric, error) {
	m := &Metric{
		Type: t.Name,
	}
	var err error
	if m.Value, err = t.Evaluate(b, id); err != nil {
		return nil, fmt.Errorf("cannot evaluate metric \"%v\" for object \"%v\" of type \"%v\"",
			t.Name, id, t.ObjectTypeName)
	}
	return m, err
}