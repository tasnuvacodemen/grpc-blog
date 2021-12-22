package main

import (
	"log"

	"github.com/tasnuvatina/grpc-blog/todo/storage/postgres"
)

func main() {
	if err := postgres.Migrate(); err != nil {
		log.Fatal(err)
	}
}
