package main

import (
	"log"

	"github.com/pasca-l/grpc-sample/server/server"
)

func main() {
	err := server.Serve("8080")
	if err != nil {
		log.Fatal(err)
	}
}
