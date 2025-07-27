package utils

import (
	"context"
	"log/slog"
	"main/src/types"
	"time"
)

// MongoClient is the interface for Mongo actions that need to be performed
type MongoClient interface {
	InsertStat(stat types.DataPoint) error
	GetStats(query string) ([]types.DataPoint, error)
}

// InsertStat inserts a single stat into the stats database
func (mongo *Mongo) InsertStat(stat types.DataPoint) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := mongo.Client.Database("plantTracker").Collection("stats")

	res, err := collection.InsertOne(ctx, stat)
	if err != nil {
		return err
	}
	id := res.InsertedID

	slog.Info("Successfully inserted record: %s", id)
	return nil
}

// GetStats returns all stats that match a given query
func (mongo *Mongo) GetStats(query string) ([]types.DataPoint, error) {
	return []types.DataPoint{}, nil
}
