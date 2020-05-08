package main

import (
	"context"
	"fmt"

	"github.com/danielsoro/go-mongo/database"
	"github.com/danielsoro/go-mongo/models"
	"github.com/danielsoro/go-mongo/repository"
)

func main() {
	connect := database.Connect()
	connect.Client().Ping(context.TODO(), nil)
	fmt.Println("MongoDB Connected")

	id := repository.TesteRepository{}.Insert(models.Teste{
		Value: "Aloha2",
	})

	repository.TesteRepository{}.Remove(id.Hex())
}
