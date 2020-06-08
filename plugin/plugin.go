package plugin

import (
	"fmt"
	"sync"

	"github.com/adyatlov/xp/data"
)

type Name string

var plugins = make(map[Name]Plugin)
var pluginsMu = &sync.RWMutex{}

type Plugin interface {
	Name() Name
	Description() string
	Open(url string) (data.Dataset, error)
	Compatible(url string) (bool, error)
}

func RegisterPlugin(plugin Plugin) {
	pluginsMu.Lock()
	defer pluginsMu.Unlock()
	if plugin == nil {
		panic("plugin is nil")
	}
	if _, dup := plugins[plugin.Name()]; dup {
		panic("plugin already registered: " + plugin.Name())
	}
	plugins[plugin.Name()] = plugin
}

func GetPlugin(name Name) (Plugin, error) {
	pluginsMu.RLock()
	plugin, ok := plugins[name]
	pluginsMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("unknown plugin %q (forgotten import?)", name)
	}
	return plugin, nil
}

func GetPlugins() []Plugin {
	pluginsMu.RLock()
	defer pluginsMu.RUnlock()
	res := make([]Plugin, 0, len(plugins))
	for _, plugin := range plugins {
		res = append(res, plugin)
	}
	return res
}

func GetCompatiblePlugins(url string) ([]Plugin, error) {
	pluginsMu.RLock()
	defer pluginsMu.RUnlock()
	res := make([]Plugin, 0, len(plugins))
	for _, plugin := range plugins {
		ok, err := plugin.Compatible(url)
		if err != nil {
			return nil, fmt.Errorf("cannot check compatibility of URL \"%v\" with plugin \"%v\": %w",
				url, plugin.Name(), err)
		}
		if ok {
			res = append(res, plugin)
		}
	}
	return res, nil
}
