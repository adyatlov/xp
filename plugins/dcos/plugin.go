package dcos

import (
	"fmt"
	"hash/fnv"
	"io"

	"github.com/adyatlov/xp/data"
	"github.com/mesosphere/bun/v2/bundle"
)

func init() {
	p := &data.Plugin{
		Name:        "DC/OS Bundle",
		Description: "Plugin for DC/OS diagnostics bundle",
		Open:        open,
		Compatible:  compatible,
	}
	data.RegisterPlugin(p)
}

func open(url string) (data.Dataset, error) {
	b, err := bundle.New(url)
	if err != nil {
		return nil, err
	}
	id, err := generateId(b)
	if err != nil {
		return nil, err
	}
	return &Dataset{
		id: id,
		b:  b,
	}, nil
}

func generateId(b bundle.Bundle) (data.DatasetId, error) {
	summaryReport, err := b.OpenFile("summary-report")
	if err != nil {
		return "", err
	}
	hash := fnv.New64a()
	_, err = io.Copy(hash, summaryReport)
	if err != nil {
		return "", err
	}
	return data.DatasetId(
		fmt.Sprintf("%x", hash.Sum(nil))), nil
}

func compatible(url string) (bool, error) {
	_, err := bundle.New(url)
	if err != nil {
		return false, nil
	}
	return true, nil
}
