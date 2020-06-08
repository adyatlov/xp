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
	PVTBool       PropertyValueType = 0
	PVTString                       = 1
	PVTInteger                      = 2
	PVTReal                         = 3
	PVTPercentage                   = 4
	PVTVersion                      = 5
	PVTTimestamp                    = 6
	PVTType                         = 7
	PVTFile                         = 8
	PVTObject                       = 9
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
