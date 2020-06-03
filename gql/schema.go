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
	schema.Mutation.onDatasetUpdate = schema.Subscription.DatasetsUpdated
	return schema
}

const SchemaString = `
enum PropertyValueType {
    INTEGER
    REAL
    PERCENTAGE
    VERSION
    TIMESTAMP
    TYPE
    FILE
}

schema {
    query: Query
    mutation: Mutation
    subscription: Subscription
}

type Query {
    object(datasetId: ID!, id: ID!): Object!
    datasets(Ids: [ID!]): [Dataset!]!
    plugins(url: String): [Plugin!]!
}

type Mutation {
    addDataset(plugin: String!, url: String!): Dataset!
    removeDataset(id: ID!): Boolean!
}

type Subscription {
    datasetsChanged: [Dataset!]!
}

type Object {
    type:                           ObjectType!
    id:								ID!
    name:                           String!
    children(typeNames: [String!]): [ObjectGroup!]!
    properties(typeNames: [String!]):  [Property!]!
}

type ObjectGroup {
    type: ObjectType!
    objects:  [Object!]!
    total:    Int!
}

type ObjectType {
    name:           String!
    pluralName: 	String!
    description:    String!
    properties:        [PropertyType!]!
    defaultProperties: [String!]!
}

type Property {
    type:  PropertyType!
    value: String!
}

type PropertyType {
    name:           String!
    valueType:      PropertyValueType!
    description:    String!
}

type Dataset {
    id:   ID!
    root: Object!
    plugin: Plugin!
    url: String!
    added: String!
}

type Plugin {
    name: String!
    description: String!
}
`
