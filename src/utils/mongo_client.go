package utils

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func SetupMongoClient(c *mongo.Client) *mongo.Client {
	if c != nil {
		return c
	}

	addr := os.Getenv("MONGO_URI")
	port := os.Getenv("MONGO_PORT")

	c, _ = mongo.Connect(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", addr, port)))

	return c
}
