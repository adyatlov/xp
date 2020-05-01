package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adyatlov/xp/xp"

	"github.com/adyatlov/xp/server"

	"github.com/mesosphere/bun/v2/bundle"
)

func main() {
	fmt.Println("Please connect to http://localhost:7777")
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
	e := xp.NewExplorer(&b)
	s := server.New(e)
	log.Fatal(s.Serve())
}
