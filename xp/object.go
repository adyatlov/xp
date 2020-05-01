package xp

import (
	"fmt"
	"log"
	"sort"

	"github.com/mesosphere/bun/v2/bundle"
)

type ObjectTypeName string
type ObjectId string
type ObjectName string

type ObjectType struct {
	Name              ObjectTypeName
	DisplayName       string
	PluralDisplayName string
	Description       string
	Metrics           []MetricTypeName
	DefaultMetrics    []MetricTypeName
	FindObject        func(*bundle.Bundle, ObjectId) (ObjectName, error)
	FindChildren      map[ObjectTypeName]func(*bundle.Bundle) ([]ObjectId, error)
}

type Object interface {
	TypeName() ObjectTypeName
	Id() ObjectId
	Name() ObjectName
	Metrics(...MetricTypeName) ([]*Metric, error)
	Children(...ObjectTypeName) ([]ObjectGroup, error)
	CountChildren(...ObjectTypeName) ([]ObjectGroup, error)
}

type ObjectGroup struct {
	TypeName ObjectTypeName
	Objects  []Object
	Count    int
}

func (t ObjectType) New(b *bundle.Bundle, id ObjectId) (Object, error) {
	name, err := t.FindObject(b, id)
	if err != nil {
		return nil, fmt.Errorf("cannot find object: %s", err.Error())
	}
	if id == "" {
		id = ObjectId(name)
	}
	object := &objectImpl{
		t:    t,
		id:   id,
		name: name,
	}
	return object, nil
}

func (t ObjectType) getChildren(b *bundle.Bundle,
	objectId ObjectId,
	onlyCount bool,
	childrenTypeNames []ObjectTypeName) ([]ObjectGroup, error) {
	if len(childrenTypeNames) == 0 {
		childrenTypeNames = make([]ObjectTypeName, 0, len(t.FindChildren))
		for typeName, _ := range t.FindChildren {
			childrenTypeNames = append(childrenTypeNames, typeName)
		}
	}
	if len(t.FindChildren) < len(childrenTypeNames) {
		return nil, fmt.Errorf("%v children type names specified, but only %v registered",
			len(childrenTypeNames), len(t.FindChildren))
	}
	groups := make([]ObjectGroup, 0, len(childrenTypeNames))
	for _, childrenTypeName := range childrenTypeNames {
		found := false
		for registeredTypeName, findChildren := range t.FindChildren {
			if childrenTypeName != registeredTypeName {
				continue
			}
			found = true
			childrenIds, err := findChildren()
			if err != nil {
				return nil, fmt.Errorf("error occurred when finding children of object type \"%v\"", childrenTypeName)
			}
			var group ObjectGroup
			if !onlyCount {
				children := make([]Object, 0, len(childrenIds))
				childrenType := MustGetObjectType(childrenTypeName)
				for _, childrenId := range childrenIds {
					object, err := childrenType.New(b, childrenId)
					if err != nil {
						log.Printf("cannot create object of type \"%v\" and ID %v: %v\n", childrenTypeName, childrenId, err)
						break
					}
					children = append(children, object)
				}
				group = ObjectGroup{
					TypeName: childrenTypeName,
					Objects:  children,
					Count:    len(children),
				}
			} else {
				group = ObjectGroup{
					TypeName: childrenTypeName,
					Count:    len(childrenIds),
				}
			}
			groups = append(groups, group)
			break
		}
		if !found {
			return nil, fmt.Errorf("requested children type \"%v\" doesn't exist for object type \"%v\"",
				childrenTypeName, t.Name)
		}
	}
	return groups, nil
}

func (t ObjectType) getMetrics(b *bundle.Bundle,
	id ObjectId,
	mm []MetricTypeName) ([]*Metric, error) {
	if len(mm) == 0 {
		mm = t.Metrics
	}
	metrics := make([]*Metric, 0, len(mm))
	for _, requestedMetric := range mm {
		found := false
		for _, registeredMetric := range t.Metrics {
			if requestedMetric != registeredMetric {
				continue
			}
			found = true
			metricType := MustGetMetricType(registeredMetric)
			metric, err := metricType.New(b, id)
			if err != nil {
				log.Printf("cannot create metric: \"%v\": %s\n",
					requestedMetric, err.Error())
				break
			}
			metrics = append(metrics, metric)
			break
		}
		if !found {
			return nil, fmt.Errorf("reguested metric \"%v\" doesn't exist for object type \"%v\"",
				requestedMetric, t.Name)
		}
	}
	return metrics, nil
}

func sortChildren(children []ObjectGroup) {
	types := GetObjectTypes()
	sort.Slice(children, func(i, j int) bool {
		return types[children[i].TypeName].PluralDisplayName < types[children[j].TypeName].PluralDisplayName
	})
	for _, c := range children {
		sort.Slice(c.Objects, func(i, j int) bool {
			return c.Objects[i].Name() < c.Objects[j].Name()
		})
	}
}

type objectImpl struct {
	b    *bundle.Bundle
	t    ObjectType
	id   ObjectId
	name ObjectName
}

func (o *objectImpl) TypeName() ObjectTypeName {
	return o.t.Name
}

func (o *objectImpl) Id() ObjectId {
	return o.id
}

func (o *objectImpl) Name() ObjectName {
	return o.name
}

func (o *objectImpl) Children(t ...ObjectTypeName) ([]ObjectGroup, error) {
	return o.t.getChildren(o.b, o.id, false, t)
}

func (o *objectImpl) CountChildren(t ...ObjectTypeName) ([]ObjectGroup, error) {
	return o.t.getChildren(o.b, o.id, true, t)
}

func (o *objectImpl) Metrics(mm ...MetricTypeName) ([]*Metric, error) {
	return o.t.getMetrics(o.b, o.id, mm)
}
