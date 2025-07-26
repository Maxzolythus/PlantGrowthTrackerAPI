package routes

import (
	"encoding/json"
	"log"
	"main/src/types"
	"net/http"
)

// HomeHandler adds wimsy to the system
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	response := types.SuccessResp{
		Message: "Live Long and Prosper",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatalf("Encoding Error: %s", err)
	}
}
