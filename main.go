package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
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

	conn, err := pgx.Connect(ctx, fmt.Sprintf("user=%s dbname=%s", dbUser, dbName))
	if err != nil {
		log.Fatalf("An error occured:\n%s", err)
		return;
	}
	defer conn.Close(ctx)
	// initialize Handler instance
	handler := handlers.NewHandler(conn)
	// initlaize app
	app := fiber.New(fiber.Config{
		Immutable: true,
	})
	// load routes
	routes.LoadRoutes(app, handler)
	// listen
	app.Listen(":8008")
}