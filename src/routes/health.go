package routes

import (
	"context"
	"encoding/json"
	"main/src/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

// HealthHandler displays the health of the application and its dependancies
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	statusCode := http.StatusOK
	mongoHealth := true
	mongoClient := utils.SetupMongoClient(nil)

	err := mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		statusCode = http.StatusInternalServerError
		mongoHealth = false
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]bool{"mongoHealth": mongoHealth})
}
