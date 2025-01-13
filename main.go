package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/mattg1243/willow-server/application"
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

	// connect to db
	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		log.Fatalf("An error occured:\n%s", err)
		return
	}

	defer conn.Close(ctx)

	handler := handlers.New(conn)

	app := application.New(handler)
	err = app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
