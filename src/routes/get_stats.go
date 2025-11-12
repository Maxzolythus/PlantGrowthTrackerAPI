package routes

import (
	"encoding/json"
	"log"
	"main/src/types"
	"main/src/utils"
	"net/http"
)

// GetStatsHandler queries the DB for data points
func GetStatsHandler(mongoClient utils.MongoClient) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		response := types.SuccessResp{
			Message: "Successfully Retrieved Data",
		}
		// get the values from the query

		// escape string
		//html.EscapeString()

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
	return fn
}

func getStats(mongo utils.MongoClient) ([]types.DataPoint, error) {
	return nil, nil
}
