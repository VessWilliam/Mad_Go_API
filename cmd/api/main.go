package main

import (
	"database/sql"
	"log"
	env "rest_api_gin/internal/.env"
	"rest_api_gin/internal/database"

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

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	models := database.NewModels(db)
	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "vesswilliam"),
		models:    models,
	}

	if err := serve(app); err != nil {
		log.Fatal(err)
	}
}
