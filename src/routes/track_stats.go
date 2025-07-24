package routes

import (
	"encoding/json"
	"log/slog"
	"main/src/types"
	"main/src/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// TrackStatsHandler adds provided data to the DB
func TrackStatsHandler(w http.ResponseWriter, r *http.Request) {
	// get the values from the query

	// escape string

	stat := utils.FormatStats()

	// access mongo
	mongoClient := utils.SetupMongoClient(nil)

	err := trackStats(mongoClient, stat)
	if err != nil {
		slog.Error("TrackStats Error: " + err.Error())
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "TrackStats Error: " + err.Error()})
	}

	// sucesss
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sucessfully added Stats"})
}

func trackStats(mongo *mongo.Client, stat types.DataPoint) error {
	return nil
}
