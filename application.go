package main

import (
	"context"
	"fmt"
	"log"

	configuration "github.com/danielsoro/go-mongo/config"
	"github.com/danielsoro/go-mongo/database"
)

func main() {
	config := configuration.GetConfiguation()

	// Connect to the mongodb database with config properties
	mongoDB := database.Connect(config.Database)

	// Test connection
	if err := mongoDB.Client().Ping(context.TODO(), nil); err != nil {
		log.Fatalf("connection error: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
}
