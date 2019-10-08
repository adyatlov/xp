package dcosversion

import (
	"fmt"

	"github.com/adyatlov/bunxp/explorer"

	"github.com/mesosphere/bun/v2/bundle"
)

func init() {
	var t = explorer.MetricType{
		Name:           "dcos-version",
		ObjectTypeName: "cluster",
		ValueType:      explorer.MTVersion,
		MetricName:     "DC/OS version",
		Description:    "DC/OS version installed on the cluster",
		Evaluate:       e,
	}
	explorer.RegisterMetricType(t)
}

func e(b *bundle.Bundle, id explorer.ObjectId) (interface{}, error) {
	if len(b.Hosts) == 0 {
		return nil, fmt.Errorf("cannot find a single host in the directory %s", b.Path)
	}
	var host bundle.Host
	// Obtain the version from a random host assuming that they are all the same
	for _, host = range b.Hosts {
		break
	}
	version := &struct {
		Version string
	}{}
	if err := host.ReadJSON("dcos-version", version); err != nil {
		return nil, err
	}
	return version.Version, nil
}