package routes

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"
	"main/src/types"
	"main/src/utils"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// TrackStatsHandler adds provided data to the DB
func TrackStatsHandler(w http.ResponseWriter, r *http.Request) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	response := types.SuccessResp{
		Message: "Sucessfully added Stats",
	}
	// get the values from the body
	var stat types.DataPoint
	err := json.NewDecoder(r.Body).Decode(&stat)
	if err != nil {
		SendError(w, http.StatusBadRequest, "TrackStats Error: Unable to parse request", err)
	}

	err = validate.Struct(stat)
	if err != nil {
		SendError(w, http.StatusBadRequest, "TrackStats Error: Unable to validate request", err)
	}

	// access mongo
	mongoClient := utils.SetupMongoClient(nil)

	err = trackStats(mongoClient, stat)
	if err != nil {
		SendError(w, http.StatusInternalServerError, "TrackStats Error: Unable to add stats to table", err)
	}

	// sucesss
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatalf("Encoding Error: %s", err)
	}
}

func trackStats(mongoClient *mongo.Client, stat types.DataPoint) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := mongoClient.Database("plantTracker").Collection("stats")

	res, err := collection.InsertOne(ctx, stat)
	if err != nil {
		return err
	}
	id := res.InsertedID

	slog.Info("Successfully inserted record: %s", id)

	return nil
}
