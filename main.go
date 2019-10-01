package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/adyatlov/bunxp/objects"

	_ "github.com/adyatlov/bunxp/objects/cluster"
	"github.com/mesosphere/bun/v2/bundle"
)

func main() {
	var path string
	var err error
	if len(os.Args) == 1 {
		path, err = os.Getwd()
		if err != nil {
			fmt.Printf("Error: cannot detect a working directory: %v\n", err.Error())
			os.Exit(1)
		}
	} else {
		path = os.Args[1]
	}
	b, err := bundle.NewBundle(path)
	if err != nil {
		fmt.Printf("Error: cannot create a bundle: %v\n", err.Error())
		os.Exit(1)
	}
	explorer := &objects.Explorer{
		Bundle: &b,
	}
	cluster, err := explorer.Object("cluster", "")
	if err != nil {
		fmt.Printf("Error: cannot get cluster object: %v\n", err.Error())
		os.Exit(1)
	}
	clusterPrint, err := json.Marshal(cluster)
	if err != nil {
		fmt.Printf("Error: cannot convert cluster object: %v\n", err.Error())
		os.Exit(1)
	}
	fmt.Println(string(clusterPrint))
}
