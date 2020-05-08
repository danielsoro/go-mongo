package main

import (
	"context"
	"fmt"
	"log"

	configuration "github.com/danielsoro/go-mongo/config"
	"github.com/danielsoro/go-mongo/database"
	"github.com/spf13/viper"
)

func main() {
	var config configuration.Configuration

	// Get the configuration from confi.yaml file
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("unable to load configuration file, %v", err)
	}

	if err = viper.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	// Connect to the mongodb database with config properties
	mongoDB := database.Connect(config.Database)

	// Test connection
	if err := mongoDB.Client().Ping(context.TODO(), nil); err != nil {
		log.Fatalf("connection error: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
}
