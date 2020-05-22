package main

import (
	"context"
	"fmt"

	"github.com/danielsoro/go-mongo/database"
)

func main() {
	connect := database.Connect()
	connect.Client().Ping(context.TODO(), nil)
	fmt.Println("MongoDB Connected")
}
