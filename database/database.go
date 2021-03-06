package database

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/danielsoro/go-mongo/crypto"

	configuration "github.com/danielsoro/go-mongo/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once     sync.Once
	instance *mongo.Database
)

func _isFirstParam(uri string) bool {
	return strings.Contains(uri, "?")
}

// CustomConnect returns a custom connected based in the configuration parameter
func CustomConnect(c configuration.DatabaseConfiguration) *mongo.Database {
	once.Do(func() {
		password, _ := crypto.Decrypt(c.Password)
		uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", c.Username, password, c.Host)
		firstParam := _isFirstParam(uri)

		if len(c.CaFilePath) > 0 {
			if firstParam {
				uri = fmt.Sprintf("uri"+"&tlsCAFile=%s", c.CaFilePath)
			} else {
				uri = fmt.Sprintf("uri"+"?tlsCAFile=%s", c.CaFilePath)
			}
		}

		firstParam = _isFirstParam(uri)
		if len(c.CertificateKeyFilePath) > 0 {
			if firstParam {
				uri = fmt.Sprintf("uri"+"&tlsCertificateKeyFile=%s", c.CertificateKeyFilePath)
			} else {
				uri = fmt.Sprintf("uri"+"?tlsCertificateKeyFile=%s", c.CertificateKeyFilePath)
			}
		}

		clientOptions := options.Client().ApplyURI(uri)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			panic(fmt.Errorf("connection issue: %v", err))
		}
		instance = client.Database(c.Namespace)
	})
	return instance
}

// Connect returns a connection with mongodb read default configuration from config.yml
func Connect() *mongo.Database {
	config := configuration.GetDatabaseConfiguration()
	return CustomConnect(config)
}

// Close all connection from the client
func Close() {
	if err := instance.Client().Disconnect(context.TODO()); err != nil {
		log.Fatalf("disconnect error: %v", err)
	}
}
