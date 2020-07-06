package data

import (
	"fmt"
	"time"
)

type (
	ObjectId          string
	ObjectTypeName    string
	ObjectName        string
	PropertyName      string
	PropertyValueType int
	DatasetId         string
)

const (
	PVTBool       PropertyValueType = 0 // bool
	PVTString                       = 1 // string
	PVTInteger                      = 2 // int64
	PVTReal                         = 3 // float
	PVTPercentage                   = 4 // float
	PVTVersion                      = 5 // string
	PVTTimestamp                    = 6 // time.Time
	PVTType                         = 7 // string
	PVTFile                         = 8 // string
	PVTObject                       = 9 // ObjectLink
)

type Object interface {
	// ValueType returns object type, which contains
	Type() *ObjectType
	// Id returns a unique object id. The id must be unique amongst objects of the same type.
	// The object id is not required to be unique amongst objects of different types or from
	// different datasets.
	Id() ObjectId
	// Name returns object name. Names are not required to be unique, they are supposed to be
	// human-readable and help people to distinguish one object from another.
	Name() ObjectName
	// Properties appends property values of types specified with the names argument to
	// the properties slice. The order of the properties must be the same as in the names slice.
	// If the property with a given name is not found then the correspondent value must be set to nil.
	// Properties should return all object properties if names is empty.
	// Returns error if something went wrong during the evaluation of the properties.
	Properties(properties *[]interface{}, names ...PropertyName) error
	// ChildGroup returns a group of object children of the specified type.
	// If the childType is not listed amongst the correspondent ObjectType.ChildTypes,
	// the function returns nil.
	ChildGroup(childTypeName ObjectTypeName) ObjectGroup
}

type Dataset interface {
	Id() DatasetId
	Root() (Object, error)
	Find(t ObjectTypeName, n ObjectId) (Object, error)
}

type ObjectType struct {
	Name          ObjectTypeName
	PluralName    string
	Description   string
	PropertyTypes []*PropertyType
	ChildTypes    []*ObjectType
	// below are the cache fields
	propertyTypes  map[PropertyName]*PropertyType
	propertyNames  []PropertyName
	childTypes     map[ObjectTypeName]*ObjectType
	childTypeNames []ObjectTypeName
}

// PropertyType returns nil if there is no property with the given name.
func (o *ObjectType) PropertyType(name PropertyName) *PropertyType {
	if o.propertyTypes == nil {
		o.propertyTypes = make(map[PropertyName]*PropertyType, len(o.PropertyTypes))
		for _, t := range o.PropertyTypes {
			o.propertyTypes[t.Name] = t
		}
	}
	return o.propertyTypes[name]
}

func (o *ObjectType) PropertyNames() []PropertyName {
	if o.propertyNames == nil {
		o.propertyNames = make([]PropertyName, 0, len(o.PropertyTypes))
		for _, t := range o.PropertyTypes {
			o.propertyNames = append(o.propertyNames, t.Name)
		}
	}
	return o.propertyNames
}

func (o *ObjectType) ChildType(name ObjectTypeName) *ObjectType {
	if o.childTypes == nil {
		o.childTypes = make(map[ObjectTypeName]*ObjectType, len(o.ChildTypes))
		for _, t := range o.ChildTypes {
			o.childTypes[t.Name] = t
		}
	}
	return o.childTypes[name]
}

func (o *ObjectType) ChildTypeNames() []ObjectTypeName {
	if o.childTypeNames == nil {
		o.childTypeNames = make([]ObjectTypeName, 0, len(o.ChildTypes))
		for _, t := range o.ChildTypes {
			o.childTypeNames = append(o.childTypeNames, t.Name)
		}
	}
	return o.childTypeNames
}

type PropertyType struct {
	Name        PropertyName
	ValueType   PropertyValueType
	Description string
}

type ObjectGroup interface {
	Type() *ObjectType
	// All appends all objects to the given array. Use it when you expect only a few objects.
	// Returns error if it cannot retrieve objects due to some technical problems.
	All(*[]Object) error
	// TotalCount returns the amount of objects.
	TotalCount() int
	// Pager should return nil, if the pager is not implemented.
	// Implement pager if you expect many objects in the group.
	// NOTE: don't implement it so far. The design is not finished. Use All instead so far.
	Pager() ObjectPager
}

type ObjectPager interface {
	// Page appends objects returned by the query to the objects argument:
	//
	// sortBy sets the sort order in which all the objects in the correspondent group should be
	// sorted before slicing them. If the sortBy parameter is an empty string or a non-existing
	// property name, pager should sort objects by their names.
	//
	// first is the maximum number of the objects to append.
	//
	// after is the id of the object in the sorted group after which pager returns first objects.
	//
	// The hasMore return value indicates if there is more objects in the sorted object group
	// after the first first objects after after.
	// Returns error if it cannot retrieve objects due to some technical problems.
	Page(objects *[]Object,
		sortBy PropertyName,
		first int,
		after ObjectId) (hasMore bool, err error)
}

type ObjectLink struct {
	Type ObjectTypeName
	Id   ObjectId
}

func ValidatePropertyValue(t PropertyValueType, v interface{}) error {
	var ok bool
	expect := ""
	switch t {
	case PVTBool:
		_, ok = v.(bool)
		expect = "bool"
	case PVTString:
		_, ok = v.(string)
		expect = "string"
	case PVTInteger:
		_, ok = v.(int)
		expect = "int"
	case PVTReal:
		_, ok = v.(float64)
		expect = "float64"
	case PVTPercentage:
		_, ok = v.(float64)
		expect = "float64"
	case PVTVersion:
		_, ok = v.(string)
		expect = "string"
	case PVTTimestamp:
		_, ok = v.(time.Time)
		expect = "time.Time"
	case PVTType:
		_, ok = v.(string)
		expect = "string"
	case PVTFile:
		_, ok = v.(string)
		expect = "string"
	case PVTObject:
		_, ok = v.(ObjectLink)
		expect = "data.ObjectLink"
	}
	if !ok {
		return fmt.Errorf(" Expected property type %s, got %T", expect, v)
	}
	return nil
}
