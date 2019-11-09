package xp

import (
	"fmt"

	"github.com/mesosphere/bun/v2/bundle"
)

type MetricValueType string
type MetricTypeName string

const (
	MTInteger    MetricValueType = "integer"
	MTReal                       = "real"
	MTPercentage                 = "percentage"
	MTVersion                    = "version"
	MTTimestamp                  = "timestamp"
	MTType                       = "type"
)

type Metric struct {
	Type  MetricTypeName
	Value interface{}
}

type MetricType struct {
	Name           MetricTypeName
	ObjectTypeName ObjectTypeName
	ValueType      MetricValueType
	DisplayName    string
	Description    string
	Evaluate       func(*bundle.Bundle, ObjectId) (interface{}, error)
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
