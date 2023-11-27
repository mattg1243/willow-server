package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/mattg1243/sqlc-fiber/handlers"
	"github.com/mattg1243/sqlc-fiber/routes"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=postgres dbname=postgres ")
	if err != nil {
		fmt.Printf("An error occured:\n%s", err)
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