// Package mem provides in-memory implementation of the interfaces declared in the
//	github.com/adyatlov/xp/data package.
// Use it when a dataset that you are going to expose is relatively small.
package mem

import (
	"fmt"

	"github.com/adyatlov/xp/data"
)

// Check implementation
var _ data.Object = &Object{}
var _ data.ObjectGroup = &ObjectGroup{}
var _ data.Dataset = &Dataset{}

type Object struct {
	t              *data.ObjectType
	id             data.ObjectId
	name           data.ObjectName
	properties     map[data.PropertyName]interface{}
	childrenGroups map[data.ObjectTypeName]map[data.ObjectId]data.Object
}

func newObject(t *data.ObjectType, id data.ObjectId, name data.ObjectName) *Object {
	object := &Object{t: t, id: id, name: name}
	object.properties = make(map[data.PropertyName]interface{})
	object.childrenGroups = make(map[data.ObjectTypeName]map[data.ObjectId]data.Object)
	return object
}

func (o *Object) AddProperty(n data.PropertyName, p interface{}) {
	o.assertPropertyCompatible(n, p)
	o.properties[n] = p
}

func (o *Object) assertPropertyCompatible(name data.PropertyName, p interface{}) {
	if _, ok := o.properties[name]; ok {
		panic(fmt.Sprintf("object %q of type %q already has property %q",
			o.id, o.t.Name, name))
	}
	t := o.Type().PropertyType(name)
	if t == nil {
		panic(fmt.Sprintf("object %q of type %q cannot have property %q",
			o.id, o.t.Name, name))
	}
	if err := data.ValidatePropertyValue(t.ValueType, p); err != nil {
		panic(fmt.Sprintf("error when setting property %q of object %q of type %q: %q",
			name, o.id, o.t.Name, err))
	}
}

func (o *Object) AddChild(child data.Object) {
	if t := o.Type().ChildType(child.Type().Name); t == nil {
		panic(fmt.Sprintf("object %q of type %q cannot have child with type %q",
			o.name, o.t.Name, child.Type().Name))
	}
	if _, ok := o.childrenGroups[child.Type().Name][child.Id()]; ok {
		panic(fmt.Sprintf("object %q of type %q already has a child with id %q",
			o.id, o.t.Name, child.Id()))
	}
	if _, ok := o.childrenGroups[child.Type().Name]; !ok {
		o.childrenGroups[child.Type().Name] = make(map[data.ObjectId]data.Object)
	}
	o.childrenGroups[child.Type().Name][child.Id()] = child
}

func (o *Object) Type() *data.ObjectType {
	return o.t
}

func (o *Object) Id() data.ObjectId {
	return o.id
}

func (o *Object) Name() data.ObjectName {
	return o.name
}

func (o *Object) Properties(properties *[]interface{}, names ...data.PropertyName) error {
	if len(names) == 0 {
		names = make([]data.PropertyName, 0, len(o.t.PropertyTypes))
		for _, t := range o.t.PropertyTypes {
			names = append(names, t.Name)
		}
	}
	for _, typeName := range names {
		property := o.properties[typeName]
		*properties = append(*properties, property)
	}
	return nil
}

func (o *Object) ChildGroup(childTypeName data.ObjectTypeName) data.ObjectGroup {
	if t := o.Type().ChildType(childTypeName); t != nil {
		return ObjectGroup{parent: o, objectType: t}
	}
	return nil
}

type ObjectGroup struct {
	parent     *Object
	objectType *data.ObjectType
}

func (g ObjectGroup) children() map[data.ObjectId]data.Object {
	if t := g.parent.Type().ChildType(g.objectType.Name); t == nil {
		panic(fmt.Sprintf("illegal state: object %q of type %q cannot have child with type %q",
			g.parent.name, g.parent.Type().Name, g.objectType.Name))
	}
	return g.parent.childrenGroups[g.objectType.Name]
}

func (g ObjectGroup) Type() *data.ObjectType {
	return g.objectType
}

func (g ObjectGroup) All(objects *[]data.Object) error {
	for _, child := range g.children() {
		*objects = append(*objects, child)
	}
	return nil
}

func (g ObjectGroup) TotalCount() int {
	return len(g.children())
}

func (g ObjectGroup) Pager() data.ObjectPager {
	return nil
}

type Dataset struct {
	id      data.DatasetId
	root    data.Object
	objects map[data.ObjectTypeName]map[data.ObjectId]data.Object
}

func NewDataset(id data.DatasetId,
	rootType *data.ObjectType,
	rootId data.ObjectId,
	rootName data.ObjectName) (*Dataset, *Object) {
	dataset := &Dataset{
		id:      id,
		objects: make(map[data.ObjectTypeName]map[data.ObjectId]data.Object),
	}
	root := dataset.NewObject(rootType, rootId, rootName)
	dataset.root = root
	return dataset, root
}

func (d *Dataset) Id() data.DatasetId {
	return d.id
}

func (d *Dataset) Root() (data.Object, error) {
	return d.root, nil
}

func (d *Dataset) Find(t data.ObjectTypeName, id data.ObjectId) (data.Object, error) {
	if objects, ok := d.objects[t]; ok {
		if object, ok := objects[id]; ok {
			return object, nil
		}
	} else {
		return nil, fmt.Errorf("objects of type %v don't exist in the dataset", t)
	}
	return nil, fmt.Errorf("%v with id %v not found", t, id)
}

func (d *Dataset) NewObject(t *data.ObjectType, id data.ObjectId, name data.ObjectName) *Object {
	o := newObject(t, id, name)
	objects, ok := d.objects[o.Type().Name]
	if !ok {
		objects = make(map[data.ObjectId]data.Object)
		d.objects[o.Type().Name] = objects
	}
	objects[o.Id()] = o
	return o
}
