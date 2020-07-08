package dcos

import "github.com/adyatlov/xp/data"

var TCluster = &data.ObjectType{
	Name:        "Cluster",
	PluralName:  "Clusters",
	Description: "DC/OS cluster",
	PropertyTypes: []*data.PropertyType{
		PClusterVersion,
		PClusterVariant,
	},
	ChildTypes: []*data.ObjectType{
		TAgent,
		TFramework,
		TTask,
	},
}

var TAgent = &data.ObjectType{
	Name:        "agent",
	PluralName:  "agents",
	Description: "DC/OS cluster worker agent",
	PropertyTypes: []*data.PropertyType{
		PId,
		PAgentType,
		PAgentHostname,
	},
	ChildTypes: []*data.ObjectType{
		TFramework,
		TTask,
	},
}

var TFramework = &data.ObjectType{
	Name:        "framework",
	PluralName:  "frameworks",
	Description: "Mesos framework",
	PropertyTypes: []*data.PropertyType{
		PId,
	},
	ChildTypes: []*data.ObjectType{
		TTask,
	},
}

var TTask = &data.ObjectType{
	Name:        "task",
	PluralName:  "tasks",
	Description: "Mesos task",
	PropertyTypes: []*data.PropertyType{
		PId,
		PFramework,
		PAgent,
	},
}

var PClusterVersion = &data.PropertyType{
	Name:        "version",
	ValueType:   data.PVTVersion,
	Description: "DC/OS version",
}

var PClusterVariant = &data.PropertyType{
	Name:        "variant",
	ValueType:   data.PVTType,
	Description: "DC/OS cluster variant",
}

var PId = &data.PropertyType{
	Name:        "id",
	ValueType:   data.PVTString,
	Description: "unique identifier",
}

var PAgentType = &data.PropertyType{
	Name:        "agent type",
	ValueType:   data.PVTType,
	Description: "DC/OS agent type",
}

var PAgentHostname = &data.PropertyType{
	Name:        "hostname",
	ValueType:   data.PVTString,
	Description: "DC/OS agent hostname",
}

var PFramework = &data.PropertyType{
	Name:        "framework",
	ValueType:   data.PVTObject,
	Description: "Mesos framework",
}

var PAgent = &data.PropertyType{
	Name:        "agent",
	ValueType:   data.PVTObject,
	Description: "DC/OS agent",
}
