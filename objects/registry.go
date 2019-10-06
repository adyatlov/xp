package objects

import (
	"fmt"
)

var objectTypes = make(map[ObjectTypeName]ObjectType)
var metricTypes = make(map[MetricTypeName]MetricType)

func RegisterObjectType(t ObjectType) {
	if _, ok := objectTypes[t.Name]; ok {
		panic(fmt.Sprintf("object type %v already registered", t.Name))
	}
	objectTypes[t.Name] = t
	for _, metricType := range t.Metrics {
		if _, ok := metricTypes[metricType.Name]; ok {
			panic(fmt.Sprintf("metric type %v already registered", metricType.Name))
		}
		metricTypes[metricType.Name] = metricType
	}
}

func GetObjectType(n ObjectTypeName) (t ObjectType, err error) {
	t, ok := objectTypes[n]
	if !ok {
		err = fmt.Errorf("there is no object of type %v", n)
	}
	return
}

func MustGetObjectType(n ObjectTypeName) ObjectType {
	t, err := GetObjectType(n)
	if err != nil {
		panic(fmt.Sprintf("Object type \"%v\" does not exist.", n))
	}
	return t
}

func GetObjectTypes() map[ObjectTypeName]ObjectType {
	return objectTypes
}

func GetMetricTypes() map[MetricTypeName]MetricType {
	return metricTypes
}
