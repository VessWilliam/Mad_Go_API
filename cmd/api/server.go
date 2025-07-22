package main

import (
	"fmt"
	"net/http"
	"os"
	"rest_api_gin/cmd/api/utils"
	"rest_api_gin/internal/handler"
	"rest_api_gin/internal/repository"
	"rest_api_gin/internal/router"
	"rest_api_gin/internal/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type application struct {
	port   int
	router *gin.Engine
}

func NewApp() (*application, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	connectString := os.Getenv("DATABASE_URL")
	PORT := utils.ParseInt(os.Getenv("PORT"), 8080)

	db, err := sqlx.Connect("postgres", connectString)
	if err != nil {
		return nil, err
	}

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandle := handler.NewUserHandler(userService)
	router := router.SetupRouter(userHandle)

	app := &application{
		port:   PORT,
		router: router,
	}

	return app, nil
}

func (a *application) Serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", a.port),
		Handler:      a.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	return srv.ListenAndServe()
}
