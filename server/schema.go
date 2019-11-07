package server

const schemaString = ` 
schema {
    query: Query
}

type Query {
    roots: [Object]
    object(objectId: String!, type: String!): Object
    objectTypes: [ObjectType!]!
    metricTypes: [MetricType!]
}

type Object {
    id:       ID!
    objectId: String!
    type:     String!
    name:     String!
    metrics:  [Metric!]
    children: [ObjectGroup!]
}

type Metric {
    type:  String!
    value: String!
}

type ObjectGroup {
    type:    String!
    objects: [Object]
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
    metricName:     String!
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
