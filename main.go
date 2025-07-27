package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"main/src/routes"
	"main/src/utils"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func main() {
	ctx := context.Background()
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	slog.Info("Performing Router Set Up....")
	r := routes.SetupRouter(nil)

	mongo := utils.NewMongoClient()
	err := mongo.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		slog.Warn("Mongo Ping Error. Ensure Mongo DB is accessable and healthy.")
	}

	slog.Info("Starting Server....")

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%s", addr, port),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  2 * time.Minute,
	}

	slog.Info("Serving")

	log.Fatal(srv.ListenAndServe())
}
