package main

import (
	"github.com/janobono/linked-in-build-a-microservice-with-go/internal/database"
	"github.com/janobono/linked-in-build-a-microservice-with-go/internal/server"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	godotenv.Load()

	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal("SERVER_PORT wrong format or not found")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatal("DB_HOST not found")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("DB_USER not found")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("DB_PASSWORD not found")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME not found")
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("DB_PORT wrong format or not found")
	}

	db, err := database.NewDatabaseClient(dbHost, dbUser, dbPassword, dbName, dbPort)
	if err != nil {
		log.Fatalf("failed to initialize Database client: %s", err)
	}
	srv := server.NewEchoServer(port, db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
