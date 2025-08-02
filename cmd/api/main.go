package main

import (
	"log"
	_ "rest_api_gin/docs" // 👈 ensure Swagger docs are registered
)

// @title           GO & Gin API
// @version         1.0
// @description     User Management System Sample App.
// @host      localhost:8080
// @BasePath  /api

func main() {
	app, err := NewApp()
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
