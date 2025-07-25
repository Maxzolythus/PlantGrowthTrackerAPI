package main

import (
	"fmt"
	"log"
	"log/slog"
	"main/src/routes"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	r := routes.SetupRouter(nil)

	// TODO: Check mongo is up before starting
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%s", addr, port),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  2 * time.Minute,
	}

	log.Fatal(srv.ListenAndServe())
}
