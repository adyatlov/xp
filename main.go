package main

import (
	"fmt"
	"log"

	"github.com/adyatlov/xp/server"
)

func main() {
	fmt.Println("Please connect to http://localhost:7777")
	s := server.New()
	log.Fatal(s.Serve())
}
