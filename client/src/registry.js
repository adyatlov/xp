const registry = {
    objectTypes: null,
    metricTypes: null
};

export default registry;

/*
objectTypes
{
  "agent": {
    "name": "agent",
    "displayName": "Agent",
    "pluralDisplayName": "Agents",
    "description": "DC/OS Agent",
    "metrics": [
      "agent-type"
    ],
    "defaultMetrics": [
      "agent-type"
    ]
  },
  "cluster": {
    "name": "cluster",
    "displayName": "Cluster",
    "pluralDisplayName": "Clusters",
    "description": "DC/OS Cluster",
    "metrics": [
      "dcos-version"
    ],
    "defaultMetrics": [
      "dcos-version"
    ]
  }
}
metricTypes
{
  "agent-type": {
    "name": "agent-type",
    "objectTypeName": "",
    "valueType": "type",
    "displayName": "Agent Type",
    "description": "Type of the DC/OS node; can be \"agent\" or \"public agent\""
  },
  "dcos-version": {
    "name": "dcos-version",
    "objectTypeName": "cluster",
    "valueType": "version",
    "displayName": "DC/OS version",
    "description": "DC/OS version installed on the cluster"
  }
}
 */
