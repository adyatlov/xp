package data

type (
	ObjectId          string
	ObjectTypeName    string
	ObjectName        string
	PropertyName      string
	PropertyValueType int
	DatasetId         string
)

const (
	PVTBool       PropertyValueType = 1
	PVTString                       = 2
	PVTInteger                      = 3
	PVTReal                         = 4
	PVTPercentage                   = 5
	PVTVersion                      = 6
	PVTTimestamp                    = 7
	PVTType                         = 8
	PVTFile                         = 9
	PVTObject                       = 10
)

type Object interface {
	Type() *ObjectType
	Id() ObjectId
	Name() ObjectName
	Properties(propertyNames ...PropertyName) ([]Property, error)
	Children(typeNames ...ObjectTypeName) ([]ObjectGroup, error)
}

type ObjectGroup interface {
	Type() *ObjectType
	Objects() []Object
	Total() int
}

type Property interface {
	Type() *PropertyType
	Value() interface{}
}

type Dataset interface {
	Id() DatasetId
	Root() (Object, error)
	Find(t ObjectTypeName, n ObjectId) (Object, error)
}

type ObjectType struct {
	Name              ObjectTypeName
	PluralName        string
	Description       string
	Properties        []*PropertyType
	DefaultProperties []*PropertyType
}

type PropertyType struct {
	Name        PropertyName
	Type        PropertyValueType
	Description string
}
