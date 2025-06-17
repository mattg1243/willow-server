package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/mattg1243/willow-server/application"
	"github.com/mattg1243/willow-server/cron"
	"github.com/mattg1243/willow-server/db"
	"github.com/mattg1243/willow-server/handlers"
)

func main() {

	ctx := context.Background()

	// load .env if in development
	prodEnv := os.Getenv("PROD")
	if prodEnv != "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if len(dbUrl) == 0 {
		log.Fatal("DATABASE_URL must be set")
	}

	clientHost := os.Getenv("CLIENT_HOST")
	if len(clientHost) == 0 {
		clientHost = "http://localhost:3000"
	}

	// run db migrations
	db.RunMigrations(dbUrl)

	// initialize conn pool
	poolConfig, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		log.Fatalf("An error occured while creating conn pool config:\n%s", err)
	}

	dbPool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatalf("An error occured while creating conn pool:\n%s", err)
	}

	defer dbPool.Close()

	handler := handlers.New(dbPool)
	// cron
	cron.StartCronJobs(db.New(dbPool))

	app := application.New(handler)
	err = app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
