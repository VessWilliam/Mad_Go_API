package main

import (
	"log"
	"os"
	"rest_api_gin/internal/seed"
	"rest_api_gin/internal/utils"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func main() {
	utils.LoadEnv()

	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration command: 'up', 'down', 'force', or 'seed'")
	}

	direction := os.Args[1]

	connectString := os.Getenv("DATABASE_URL")

	migrationPath := os.Getenv("MIGRATIONS_PATH")

	if migrationPath == "" {
		migrationPath = "file://cmd/migrate/migrations"
	}

	// Connect to database
	dbx, err := sqlx.Open("postgres", connectString)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer dbx.Close()

	driver, err := postgres.WithInstance(dbx.DB, &postgres.Config{})
	if err != nil {
		log.Fatal("Failed to create migration driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatal("Failed to initialize migrations:", err)
	}

	// Handle migration commands
	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration up failed:", err)
		}
		log.Println("Migration up completed.")

	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Migration down failed:", err)
		}
		log.Println(" Migration down completed.")

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
		log.Printf("Forced migration set to version %d\n", version)

	case "seed":
		if err := seed.Seeder(dbx); err != nil {
			log.Fatal("Seeding failed:", err)
		}
		log.Println("Seeder executed successfully.")

	default:
		log.Fatal("Invalid command. Use 'up', 'down', 'force', or 'seed'")
	}
}
