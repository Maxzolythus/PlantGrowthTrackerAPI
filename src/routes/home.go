package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

// HomeHandler adds wimsy to the system
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(map[string]string{"message": "Live Long and Prosper"})
	if err != nil {
		log.Fatalf("Encoding Error: %s", err)
	}
}
