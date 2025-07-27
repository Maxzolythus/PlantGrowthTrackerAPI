package utils

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
}

func NewMongoClient() *Mongo {
	ctx := context.Background()
	addr := os.Getenv("MONGO_URI")
	port := os.Getenv("MONGO_PORT")

	c, err := mongo.Connect(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", addr, port)))
	if err != nil {
		slog.Error(fmt.Sprintf("Error connecting to Mongo: %s", err.Error()))
	}

	defer func() {
		if err := c.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	return &Mongo{Client: c}
}
