package objects

import "github.com/mesosphere/bun/v2/bundle"

type MetricType string
type MetricName string

const (
	MTInteger    MetricType = "integer"
	MTReal                  = "real"
	MTPercentage            = "percentage"
	MTVersion               = "version"
	MTTimestamp             = "timestamp"
)

type Metric struct {
	Type        MetricType
	Name        MetricName
	Description string
	Value       interface{}
	Evaluate    func(*bundle.Bundle, *Object) (interface{}, error) `json:"-"`
}
