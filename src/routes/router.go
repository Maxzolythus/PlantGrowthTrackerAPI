package routes

import (
	"github.com/gorilla/mux"
)

func SetupRouter(r *mux.Router) *mux.Router {
	if r != nil {
		return r
	}

	r = mux.NewRouter()

	r.HandleFunc("/health", HealthHandler).Methods("GET")
	// Get Data from a picture injestion service, and add to the DB
	r.HandleFunc("/stats", TrackStatsHandler).Methods("POST")
	r.HandleFunc("/stats", GetStatsHandler).Methods("GET")
	r.HandleFunc("/stats/{id}", GetStatsHandler).Methods("DELETE")
	r.HandleFunc("/stats/{id}", GetStatsHandler).Methods("PUT")

	return r
}
