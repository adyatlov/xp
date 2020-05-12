package plugin

import (
	"fmt"

	"github.com/adyatlov/xp/data"
)

type Plugin struct {
	name        data.PluginName
	description string
	open        func(url string) (data.Dataset, error)
}

func NewPlugin(name data.PluginName, description string, open func(url string) (data.Dataset, error)) *Plugin {
	return &Plugin{name: name, description: description, open: open}
}

func (p *Plugin) Name() data.PluginName {
	return p.name
}

func (p *Plugin) Description() string {
	return p.description
}

func (p *Plugin) Open(url string) (data.Dataset, error) {
	return p.open(url)
}

func (p *Plugin) Compatible(url string) (bool, error) {
	if _, err := p.open(url); err != nil {
		return false, nil
	}
	return true, nil
}

type Dataset struct {
	id      data.DatasetId
	root    data.Object
	objects map[data.ObjectTypeName]map[data.ObjectId]data.Object
}

func NewDataset(id data.DatasetId, root data.Object) *Dataset {
	dataset := &Dataset{
		id:      id,
		root:    root,
		objects: make(map[data.ObjectTypeName]map[data.ObjectId]data.Object),
	}
	dataset.AddObject(root)
	return dataset
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
	o := NewObject(t, id, name)
	d.AddObject(o)
	return o
}

func (d *Dataset) AddObject(o data.Object) {
	objects, ok := d.objects[o.Type().Name]
	if !ok {
		objects = make(map[data.ObjectId]data.Object)
		d.objects[o.Type().Name] = objects
	}
	objects[o.Id()] = o
}

type Object struct {
	t          *data.ObjectType
	id         data.ObjectId
	name       data.ObjectName
	properties map[data.PropertyTypeName]data.Property
	children   map[data.ObjectTypeName]*ObjectGroup
}

func NewObject(t *data.ObjectType, id data.ObjectId, name data.ObjectName) *Object {
	object := &Object{t: t, id: id, name: name}
	object.properties = make(map[data.PropertyTypeName]data.Property)
	object.children = make(map[data.ObjectTypeName]*ObjectGroup)
	return object
}

func (o *Object) AddChild(child data.Object) {
	group := o.children[child.Type().Name]
	if group == nil {
		group = NewObjectGroup(child.Type())
		o.children[child.Type().Name] = group
	}
	group.AddObject(child)
}

func (o *Object) AddProperty(property data.Property) {
	if _, ok := o.properties[property.Type().Name]; ok {
		panic(fmt.Sprintf("property %v already exists in object %v",
			property.Type().Name, o.name))
	}
	o.properties[property.Type().Name] = property
}

func (o *Object) Type() *data.ObjectType {
	return o.t
}

func (o *Object) Id() data.ObjectId {
	return data.ObjectId(o.id)
}

func (o *Object) Name() data.ObjectName {
	return o.name
}

func (o Object) Children(typeNames ...data.ObjectTypeName) ([]data.ObjectGroup, error) {
	var groups []data.ObjectGroup
	if len(typeNames) == 0 {
		groups = make([]data.ObjectGroup, 0, len(o.children))
		for _, g := range o.children {
			groups = append(groups, g)
		}
		return groups, nil
	}
	groups = make([]data.ObjectGroup, 0, len(typeNames))
	for _, typeName := range typeNames {
		group, ok := o.children[typeName]
		if !ok {
			return nil, fmt.Errorf("cannot find children with typeName %v", typeName)
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (o Object) Properties(typeNames ...data.PropertyTypeName) ([]data.Property, error) {
	var properties []data.Property
	if len(typeNames) == 0 {
		properties = make([]data.Property, 0, len(o.properties))
		for _, property := range o.properties {
			properties = append(properties, property)
		}
		return properties, nil
	}
	properties = make([]data.Property, 0, len(typeNames))
	for _, typeName := range typeNames {
		property, ok := o.properties[typeName]
		if !ok {
			return nil, fmt.Errorf("cannot find property with typeName %v", typeName)
		}
		properties = append(properties, property)
	}
	return properties, nil
}

type Property struct {
	t     *data.PropertyType
	value interface{}
}

func NewProperty(t *data.PropertyType, value interface{}) *Property {
	return &Property{t: t, value: value}
}

func (p *Property) Type() *data.PropertyType {
	return p.t
}

func (p *Property) Value() interface{} {
	return p.value
}

type ObjectGroup struct {
	objects    []data.Object
	objectType *data.ObjectType
}

func NewObjectGroup(objectType *data.ObjectType) *ObjectGroup {
	group := &ObjectGroup{objectType: objectType}
	group.objects = make([]data.Object, 0)
	return group
}

func (o *ObjectGroup) AddObject(object data.Object) {
	if o.objectType != object.Type() {
		panic(fmt.Sprintf("wrong object type, expected %v got %v",
			o.objectType, object.Type()))
	}
	o.objects = append(o.objects, object)
}

func (o *ObjectGroup) Type() *data.ObjectType {
	return o.objectType
}

func (o *ObjectGroup) Objects() []data.Object {
	return o.objects
}

func (o *ObjectGroup) Total() int {
	return len(o.objects)
}
