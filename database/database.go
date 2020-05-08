package database

import (
	"context"
	"fmt"
	"log"
	"sync"

	configuration "github.com/danielsoro/go-mongo/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once     sync.Once
	instance *mongo.Database
)

// Connect returns a connection with mongodb
func Connect(config configuration.DatabaseConfiguration) *mongo.Database {
	once.Do(func() {
		clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s", config.Username, config.Password, config.Host))
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			panic(fmt.Errorf("Conection issue: %v", err))
		}
		instance = client.Database("test")
	})
	return instance
}

// Close all connection from the client
func Close() {
	if err := instance.Client().Disconnect(context.TODO()); err != nil {
		log.Fatalf("disconnect error: %v", err)
	}
}