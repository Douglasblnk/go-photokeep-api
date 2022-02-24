package main

import (
	graphqlserver "go-dynamodb/api/graph"
	"go-dynamodb/internal/database"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	err = database.LoadAWS()

	if err != nil {
		log.Fatal("cannot load AWS env:", err)
	}
}

func main() {
	graphqlserver.Start()
}
