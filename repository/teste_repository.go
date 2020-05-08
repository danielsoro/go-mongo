package repository

import (
	"context"
	"log"

	"github.com/danielsoro/go-mongo/database"
	"github.com/danielsoro/go-mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	collection = database.Connect().Collection(models.TesteCollection)
)

type TesteRepository struct {
}

func (tr TesteRepository) Insert(t models.Teste) primitive.ObjectID {
	result, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
		log.Fatalf("error to insert: %v", t)
	}
	return result.InsertedID.(primitive.ObjectID)
}

func (tr TesteRepository) Remove(id string) bool {
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}
	one, err := collection.DeleteOne(context.TODO(), bson.M{"_id": hex})
	if err != nil {
		log.Fatalf("erro to delete: %s", id)
	}
	return one.DeletedCount > 0
}
