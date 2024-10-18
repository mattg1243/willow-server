package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mattg1243/willow-server/handlers"
)

type App struct {
	router http.Handler
}

func New(h *handlers.Handler) *App {
	app := &App{ 
		router: loadRoutes(h),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr: ":8008",
		Handler: a.router,
	}
	fmt.Println("Server listening...")
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start the server: %w", err)
	}

	return nil
}