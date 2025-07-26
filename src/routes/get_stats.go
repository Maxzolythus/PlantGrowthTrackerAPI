package routes

import (
	"encoding/json"
	"log"
	"main/src/types"
	"main/src/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// GetStatsHandler queries the DB for data points
func GetStatsHandler(w http.ResponseWriter, r *http.Request) {
	response := types.SuccessResp{
		Message: "Successfully Retrieved Data",
	}
	// get the values from the query

	// escape string
	//html.EscapeString()

	// access mongo or w/e and shove it in
	mongoClient := utils.SetupMongoClient(nil)

	dataPoints, err := getStats(mongoClient)
	if err != nil {
		SendError(w, http.StatusInternalServerError, "GetStats Error: Unable to retrieve stats", err)
	}

	response.Data = dataPoints

	// sucesss
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatalf("Encoding Error: %s", err)
	}
}

func getStats(mongo *mongo.Client) ([]types.DataPoint, error) {
	return nil, nil
}
