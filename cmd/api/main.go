package main

import (
	"database/sql"
	"log"
	env "rest_api_gin/Internal/.env"
	"rest_api_gin/Internal/database"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	models := database.NewModels(db)
	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "gingingo"),
		models:    models,
	}

	if err := serve(app); err != nil {
		log.Fatal(err)
	}
}
