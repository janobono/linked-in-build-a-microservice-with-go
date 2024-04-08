package main

import (
	"github.com/janobono/linked-in-build-a-microservice-with-go/internal/database"
	"github.com/janobono/linked-in-build-a-microservice-with-go/internal/server"
	"log"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to initialize Database client: %s", err)
	}
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
