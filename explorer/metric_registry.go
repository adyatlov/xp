package explorer

import (
	"fmt"
)

var metricTypes = make(map[MetricTypeName]MetricType)

func RegisterMetricType(t MetricType) {
	if _, ok := metricTypes[t.Name]; ok {
		panic(fmt.Sprintf("metric type %v already registered", t.Name))
	}
	metricTypes[t.Name] = t
}

func GetMetricType(n MetricTypeName) (t MetricType, err error) {
	t, ok := metricTypes[n]
	if !ok {
		err = fmt.Errorf("there is no metric of type %v", n)
	}
	return
}

func GetMetricTypes() map[MetricTypeName]MetricType {
	return metricTypes
}

func MustGetMetricType(n MetricTypeName) MetricType {
	t, err := GetMetricType(n)
	if err != nil {
		panic(fmt.Sprintf("Metric type \"%v\" does not exist.", n))
	}
	return t
}
