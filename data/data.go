package data

type (
	ObjectId          string
	ObjectTypeName    string
	ObjectName        string
	PropertyTypeName  string
	PropertyValueType string
	DatasetId         string
	PluginName        string
)

const (
	MVTBool       PropertyValueType = "bool"
	MVTString     PropertyValueType = "string"
	MVTInteger    PropertyValueType = "integer"
	MVTReal       PropertyValueType = "real"
	MVTPercentage PropertyValueType = "percentage"
	MVTVersion    PropertyValueType = "version"
	MVTTimestamp  PropertyValueType = "timestamp"
	MVTType       PropertyValueType = "type"
	MVTFile       PropertyValueType = "file"
	MVTObject     PropertyValueType = "object"
)

type Object interface {
	Type() *ObjectType
	Id() ObjectId
	Name() ObjectName
	Children(typeNames ...ObjectTypeName) ([]ObjectGroup, error)
	Properties(typeNames ...PropertyTypeName) ([]Property, error)
}

type ObjectGroup interface {
	Type() *ObjectType
	Objects() []Object
	Total() int
}

type ObjectType struct {
	Name              ObjectTypeName
	PluralName        string
	Description       string
	Properties        []*PropertyType
	DefaultProperties []*PropertyType
}

type Property interface {
	Type() *PropertyType
	Value() interface{}
}

type PropertyType struct {
	Name        PropertyTypeName
	ValueType   PropertyValueType
	Description string
}

type Dataset interface {
	Id() DatasetId
	Root() (Object, error)
	Find(t ObjectTypeName, n ObjectId) (Object, error)
}

type Plugin interface {
	Name() PluginName
	Description() string
	Open(url string) (Dataset, error)
	Compatible(url string) (bool, error)
}
