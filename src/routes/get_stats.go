package routes

import (
	"encoding/json"
	"log/slog"
	"main/src/types"
	"main/src/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// GetStatsHandler queries the DB for data points
func GetStatsHandler(w http.ResponseWriter, r *http.Request) {
	// get the values from the query

	// escape string
	//html.EscapeString()

	// access mongo or w/e and shove it in
	mongoClient := utils.SetupMongoClient(nil)

	dataPoints, err := getStats(mongoClient)
	if err != nil {
		slog.Error("GetStats Error: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "GetStats Error: " + err.Error()})
	}

	// sucesss
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]types.DataPoint{"message": dataPoints})
}

func getStats(mongo *mongo.Client) ([]types.DataPoint, error) {
	return nil, nil
}
