package main

import (
	"log"

	"github.com/pasca-l/grpc-sample/client/client"
)

func main() {
	// host name should match name from `docker ps`
	err := client.Client("grpc-sample-server-1:8080")
	if err != nil {
		log.Fatal(err)
	}
}
