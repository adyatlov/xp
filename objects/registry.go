package objects

import (
	"fmt"
)

var objectTypeRegistry = make(map[ObjectTypeName]*ObjectType)

func RegisterObjectType(t *ObjectType) {
	if _, ok := objectTypeRegistry[t.Name]; ok {
		panic(fmt.Sprintf("object type %v already registered", t.Name))
	}
	objectTypeRegistry[t.Name] = t
}

func GetObjectType(n ObjectTypeName) (t *ObjectType, err error) {
	t, ok := objectTypeRegistry[n]
	if !ok {
		err = fmt.Errorf("there is no object of type %v", n)
	}
	return
}
