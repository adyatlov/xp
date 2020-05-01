package server

const schemaString = ` 
schema {
    query: Query
}

type Query {
    root: Object!
    object(typeName: String!, objectId: String!): Object
    objectTypes: [ObjectType!]!
    metricTypes: [MetricType!]
}

type Object {
    id:                             ID!
    typeName:                       String!
    objectId:                       String!
    name:                           String!
    metrics(names: [String!]):      [Metric!]
    children(typeNames: [String!]): [ObjectGroup!]
    childrenCount(typeNames: [String!]): [ObjectGroup!]
}

type Metric {
    typeName:  String!
    value: String!
}

type ObjectGroup {
    typeName:    String!
    objects: [Object!]
    count: Int!
}

type ObjectType {
    name:              String!
    displayName:       String!
    pluralDisplayName: String!
    description:       String!
    metrics:           [String!]
    defaultMetrics:    [String!]
}

type MetricType {
    name:           String!
    objectTypeName: String!
    valueType:      MetricValueType!
    displayName:    String!
    description:    String!
}

enum MetricValueType {
    INTEGER
    REAL
    PERCENTAGE
    VERSION
    TIMESTAMP
    TYPE
}
`
