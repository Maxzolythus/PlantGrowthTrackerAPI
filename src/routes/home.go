package routes

import (
	"encoding/json"
	"net/http"
)

// HomeHandler adds wimsy to the system
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{"message": "Live Long and Prosper"})
}
