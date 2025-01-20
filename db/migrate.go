package db

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func RunMigrations(dbUrl string) {
	m, err := migrate.New("file://db/migrations", dbUrl)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	log.Println("Running migrations..")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
	log.Println("Migrations complete.")
} 