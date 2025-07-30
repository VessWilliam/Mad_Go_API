package main

import (
	"log"
	"os"
	"rest_api_gin/internal/seed"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please Provide a migration direction: 'up' or 'down'")
	}

	direction := os.Args[1]

	err := godotenv.Load()
	if err != nil {
		return
	}

	connectString := os.Getenv("DATABASE_URL")

	dbx, err := sqlx.Open("postgres", connectString)
	if err != nil {
		log.Fatal(err)
	}

	defer dbx.Close()

	driver, err := postgres.WithInstance(dbx.DB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}

	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	case "force":
		if len(os.Args) < 3 {
			log.Fatal("Please provide a version number for 'force'")
		}
		version, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid version number: %v", err)
		}
		if err := m.Force(version); err != nil {
			log.Fatalf("Force migration failed: %v", err)
		}
	case "seed":
		if err := seed.Seeder(dbx); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Invalid direction. Use 'Up' or 'Down'")

	}

}
