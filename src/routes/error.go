package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"main/src/types"
	"net/http"
)

// SendError handles errors by sending a standardized error message
func SendError(w http.ResponseWriter, status int, message string, err error) {
	response := types.ErrorResp{
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
		slog.Error(fmt.Sprintf("%s : %s", response.Message, response.Error))
	} else {
		slog.Error(response.Message)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatalf("Encoding Error: %s", err)
	}
}
