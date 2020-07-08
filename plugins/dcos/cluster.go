package dcos

import (
	"github.com/adyatlov/xp/data"
	"github.com/mesosphere/bun/v2/bundle"
)

type cluster struct {
	id   data.ObjectId
	name data.ObjectName
	b    bundle.Bundle
}

func newCluster(b bundle.Bundle) (*cluster, error) {
	name, err := getClusterName(b)
	if err != nil {
		return nil, err
	}
	c := cluster{
		id:   data.ObjectId(name),
		name: data.ObjectName(name),
		b:    b,
	}
	return &c, nil
}

func (c *cluster) Type() *data.ObjectType {
	return TCluster
}

func (c *cluster) Id() data.ObjectId {
	return c.id
}

func (c *cluster) Name() data.ObjectName {
	return c.name
}

func (c *cluster) Properties(properties *[]interface{}, names ...data.PropertyName) error {
	for _, name := range names {
		switch name {
		case PClusterVersion.Name:
			*properties = append(*properties, c.version())
		case PClusterVariant.Name:
			*properties = append(*properties, c.variant())
		default:
			*properties = append(*properties, nil)
		}

	}
	return nil
}

func (c *cluster) ChildGroup(childTypeName data.ObjectTypeName) data.ObjectGroup {
	switch childTypeName {
	case TAgent.Name:
		return newAgentGroup(c.b)
	}
	return nil
}

func getClusterName(b bundle.Bundle) (string, error) {
	config := struct {
		ClusterName string `json:"cluster_name"`
	}{}
	b.ForEachDirectory("expanded-config",
		func(d bundle.Directory) bool {
			if err := d.ReadJSON("expanded-config", &config); err != nil {
				return false
			}
			return true
		})
	return config.ClusterName, nil
}

func (c *cluster) version() interface{} {
	v := struct {
		Version string `json:"version"`
	}{}
	f := func(d bundle.Directory) bool {
		if err := d.ReadJSON("dcos-version", &v); err != nil {
			return false
		}
		return true
	}
	c.b.ForEachDirectory("dcos-version", f)
	return v.Version
}

func (c *cluster) variant() interface{} {
	v := struct {
		Variant string `json:"dcos_variant"`
	}{}
	f := func(d bundle.Directory) bool {
		if err := d.ReadJSON("expanded-config", &v); err != nil {
			return false
		}
		return true
	}
	c.b.ForEachDirectory("expanded-config", f)
	return v.Variant
}
