package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rest_api_gin/internal/handler"
	"rest_api_gin/internal/repository"
	"rest_api_gin/internal/router"
	"rest_api_gin/internal/service"
	"strconv"
	"time"

	_ "rest_api_gin/docs"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
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

	PORT, _ := strconv.Atoi(os.Getenv("PORT"))
	connectString := os.Getenv("DATABASE_URL")

	db, err := sqlx.Connect("postgres", connectString)
	if err != nil {
		return nil, err
	}

	// Register User route & Role route
	roleRepo := repository.NewRolesRepo(db)
	userRepo := repository.NewUserRepo(db)

	// Inject both into service
	userService := service.NewUserService(userRepo, roleRepo)
	roleService := service.NewRolesService(roleRepo)

	userHandle := handler.NewUserHandler(userService)
	roleHandle := handler.NewRoleHandle(roleService)

	//Setup router
	router := router.SetupRouter(userHandle, roleHandle)

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
	log.Printf("Server listening Port : %d\n", a.port)
	return srv.ListenAndServe()
}
