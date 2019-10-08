package explorer

import (
	"fmt"
)

var objectTypes = make(map[ObjectTypeName]ObjectType)

func RegisterObjectType(t ObjectType) {
	for _, dm := range t.DefaultMetrics {
		ok := false
		for _, m := range t.Metrics {
			if dm == m {
				ok = true
			}
		}
		if !ok {
			panic(fmt.Sprintf("Default metrics %v is not listed amongst all metrics", dm))
		}
	}
	if _, ok := objectTypes[t.Name]; ok {
		panic(fmt.Sprintf("object type %v already registered", t.Name))
	}
	objectTypes[t.Name] = t
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
