package main

import (
	"log"
	env "rest_api_gin/internal/.env"
	"rest_api_gin/internal/database"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	dsn := "postgres://postgres:root123@localhost:5433/madevent?sslmode=disable"

	dbx, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer dbx.Close()

	models := database.NewModels(dbx)
	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "vesswilliam"),
		models:    models,
	}

	if err := serve(app); err != nil {
		log.Fatal(err)
	}
}
