package dcos

import (
	"github.com/adyatlov/xp/data"
	"github.com/mesosphere/bun/v2/bundle"
)

type framework struct {
	id     string
	name   string
	active bool
}

func (f framework) Type() *data.ObjectType {
	return TFramework
}

func (f framework) Id() data.ObjectId {
	return data.ObjectId(f.id)
}

func (f framework) Name() data.ObjectName {
	return data.ObjectName(f.name)
}

func (f framework) Properties(properties *[]interface{}, names ...data.PropertyName) error {
	for _, name := range names {
		switch name {
		case PId.Name:
			*properties = append(*properties, f.id)
		case PFrameworkActive.Name:
			*properties = append(*properties, f.active)
		default:
			*properties = append(*properties, nil)
		}
	}
	return nil
}

func (f framework) ChildGroup(childTypeName data.ObjectTypeName) data.ObjectGroup {
	return nil
}

type frameworkGroup struct {
	frameworks []*framework
}

func newFrameworkGroup(b bundle.Bundle) *frameworkGroup {
	g := frameworkGroup{}
	ff := struct {
		Frameworks []struct {
			Id     string
			Name   string
			Active bool
		}
	}{}
	do := func(d bundle.Directory) bool {
		err := d.ReadJSON("mesos-master-frameworks", &ff)
		if err != nil {
			return false
		}
		return true
	}
	b.ForEachDirectory("mesos-master-frameworks", do)
	g.frameworks = make([]*framework, 0, len(ff.Frameworks))
	for _, f := range ff.Frameworks {
		g.frameworks = append(g.frameworks, &framework{
			id:     f.Id,
			name:   f.Name,
			active: f.Active,
		})
	}
	return &g
}

func (g frameworkGroup) Type() *data.ObjectType {
	return TFramework
}

func (g frameworkGroup) All(frameworks *[]data.Object) error {
	for _, f := range g.frameworks {
		*frameworks = append(*frameworks, f)
	}
	return nil
}

func (g frameworkGroup) TotalCount() int {
	return len(g.frameworks)
}

func (g frameworkGroup) Pager() data.ObjectPager {
	return nil
}
