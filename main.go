package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/mattg1243/sqlc-fiber/handlers"
	"github.com/mattg1243/sqlc-fiber/routes"
)

func main() {
	ctx := context.Background()

	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")

	var conn *pgx.Conn

	// if we are running in docker, we need to use the docker host
	if os.Getenv("DOCKER") == "true" {
		dbHost := os.Getenv("DB_HOST")
		dbPass := os.Getenv("DB_PASSWORD")
		conn, err = pgx.Connect(ctx, fmt.Sprintf("host=%s user=%s password=%s dbname=%s", dbHost, dbUser, dbPass, dbName))
		if err != nil {
			log.Fatalf("An error occured:\n%s", err)
			return
		}
	} else {
		conn, err = pgx.Connect(ctx, fmt.Sprintf("user=%s dbname=%s", dbUser, dbName))
		if err != nil {
			log.Fatalf("An error occured:\n%s", err)
			return
		}
	}

	defer conn.Close(ctx)
	// initialize Handler instance
	handler := handlers.NewHandler(conn)
	// initlaize app
	app := fiber.New(fiber.Config{
		Immutable: true,
	})
	app.Use(logger.New())
	// load routes
	routes.LoadRoutes(app, handler)
	// listen
	app.Listen(":8008")
}
