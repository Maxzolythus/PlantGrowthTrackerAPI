package routes

import (
	"encoding/json"
	"log"
	"main/src/types"
	"main/src/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// TrackStatsHandler adds provided data to the DB
func TrackStatsHandler(w http.ResponseWriter, r *http.Request) {
	mongoClient := utils.NewMongoClient()

	// get the values from the body
	var stat types.DataPoint
	err := json.NewDecoder(r.Body).Decode(&stat)
	if err != nil {
		SendError(w, http.StatusBadRequest, "TrackStats Error: Unable to parse request", err)
	}

	statusCode, message, err := trackStats(mongoClient, stat)
	if err != nil {
		SendError(w, statusCode, message, err)
	}

	// sucesss
	w.WriteHeader(statusCode)
	err = json.NewEncoder(w).Encode(types.SuccessResp{
		Message: message,
	})
	if err != nil {
		log.Fatalf("Encoding Error: %s", err)
	}
}

func trackStats(mongo utils.MongoClient, stat types.DataPoint) (int, string, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(stat)
	if err != nil {
		return http.StatusBadRequest, "TrackStats Error: Unable to validate request", err
	}

	err = mongo.InsertStat(stat)
	if err != nil {
		return http.StatusInternalServerError, "TrackStats Error: Unable add stat to database", err
	}

	return http.StatusOK, "Sucessfully added Stats", nil
}
