package routes

import (
	"encoding/json"
	"net/http"
)

// HealthHandler displays the health of the application and its dependancies
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// TODO: Check the status of our DB and include them in the response.
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
