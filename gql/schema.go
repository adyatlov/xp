package gql

type Schema struct {
	Query
	Mutation
	Subscription
	datasets *datasetRegistry
}

func NewSchema() *Schema {
	schema := &Schema{}
	schema.datasets = NewDatasetRegistry()
	schema.Query.datasets = schema.datasets
	schema.Subscription = newSubscription(schema.datasets)
	schema.Mutation.datasets = schema.datasets
	schema.Mutation.onDatasetAdded = schema.Subscription.NotifyDatasetAdded
	schema.Mutation.onDatasetRemoved = schema.Subscription.NotifyDatasetRemoved
	return schema
}

const SchemaString = `
enum PropertyValueType {
    BOOL
    STRING
    INTEGER
    REAL
    PERCENTAGE
    VERSION
    TIMESTAMP
    TYPE
    FILE
    OBJECT
}

schema {
    query: Query
    mutation: Mutation
    subscription: Subscription
}

type Query {
    node(id: ID):                   Node
    datasets:                       [Dataset!]
    plugins:                        [Plugin!]
    compatiblePlugins(url: String): [Plugin!]
}

type Mutation {
    addDataset(pluginName: String!, url: String!): Dataset!
    removeDataset(id: ID!):                        Boolean!
}

type Subscription {
    datasetUpdated: DatasetEvent!
}

type DatasetEvent {
    eventType: String!
    idToRemove: ID
    dataset: Dataset
}

interface Node {
    id: ID!
}

type Object implements Node {
    id:   ID!
    type: ObjectType!
    name: String!
    properties(
        typeNames: [String!]
        first: Int
        after: ID
    ): PropertiesConnection
    childGroup(typeName:  String): ChildGroup
    childGroups(typeNames:  [String!]): [ChildGroup]
    firstAvailableChildGroupTypeName: String
}

type Property implements Node {
    id:    ID!
    type:  PropertyType!
    value: String!
}

type PropertiesConnection {
    totalCount: Int!
    edges: [PropertyEdge]
    pageInfo: PageInfo!
}

type PropertyEdge {
    cursor: ID!
    node: Property
}

type ChildGroup {
    id:    ID!
    type: ObjectType!
    totalCount: Int!
    children(
        first:      String
        after:      ID
    ): ChildrenConnection
}

type ChildrenConnection {
    totalCount: Int!
    edges: [ChildEdge]
    pageInfo: PageInfo!
}

type ChildEdge {
    cursor: ID!
    node:   Object
}

type PageInfo {
    startCursor: ID!
    endCursor:   ID!
    hasNextPage: Boolean!
}

type Dataset implements Node {
    id:     ID!
    root:   Object!
    plugin: Plugin!
    url:    String!
    added:  String!
}

type Plugin implements Node {
    id:          ID!
    name:        String!
    description: String!
}

type ObjectType {
    name:        String!
    pluralName:  String!
    description: String!
    propertyTypes:  [PropertyType!]!
}

type PropertyType {
    name:        String!
    valueType:   PropertyValueType!
    description: String!
}
`
