package main

import (
	"log"
	_ "rest_api_gin/docs" // ensure Swagger docs are registered
)

// @title           GO & Gin API
// @version         1.0
// @description     User Management System Sample App.
// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Provide your access token in the Authorization header using the Bearer scheme.

func main() {
	app, err := NewApp()
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
