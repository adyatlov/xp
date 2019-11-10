package server

const schemaString = ` 
schema {
    query: Query
}

type Query {
    root: Object!
    object(type: String!, objectId: String!): Object
    objectTypes: [ObjectType!]!
    metricTypes: [MetricType!]
}

type Object {
    id:                        ID!
    type:                      String!
    objectId:                  String!
    name:                      String!
    metrics(names: [String!]!):  [Metric!]
    children:                  [ObjectGroup!]
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
    displayName:     String!
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
