package main

import (
	"github.com/jmichiels/cloud-functions-tests/cmd/demo-client/cmds"
	"log"
)

func main() {
	if err := cmds.Execute(); err != nil {
		log.Fatal(err)
	}
}
