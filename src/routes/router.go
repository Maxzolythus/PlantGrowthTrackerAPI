package routes

import (
	"main/src/utils"

	"github.com/gorilla/mux"
)

func SetupRouter(r *mux.Router, mongo utils.MongoClient) *mux.Router {
	if r != nil {
		return r
	}

	r = mux.NewRouter()

	r.HandleFunc("/health", HealthHandler).Methods("GET")
	// Get Data from a picture injestion service, and add to the DB
	r.HandleFunc("/stats", TrackStatsHandler(mongo)).Methods("POST")
	r.HandleFunc("/stats", GetStatsHandler(mongo)).Methods("GET")
	r.HandleFunc("/stats/{id}", GetStatsHandler(mongo)).Methods("DELETE")
	r.HandleFunc("/stats/{id}", GetStatsHandler(mongo)).Methods("PUT")

	return r
}
