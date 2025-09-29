package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rest_api_gin/internal/handler"
	"rest_api_gin/internal/middleware"
	"rest_api_gin/internal/repository"
	"rest_api_gin/internal/router"
	"rest_api_gin/internal/service"
	"rest_api_gin/internal/utils"
	"strconv"
	"time"

	_ "rest_api_gin/docs"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type application struct {
	port   int
	router *gin.Engine
}

func NewApp() (*application, error) {

	utils.LoadEnv()

	PORT, _ := strconv.Atoi(os.Getenv("PORT"))
	connectString := os.Getenv("DATABASE_URL")

	db, err := sqlx.Connect("postgres", connectString)
	if err != nil {
		return nil, err
	}

	// Repos
	roleRepo := repository.NewRolesRepo(db)
	userRepo := repository.NewUserRepo(db)

	// Services
	userService := service.NewUserService(userRepo, roleRepo)
	roleService := service.NewRolesService(roleRepo)

	// JWT service
	secret := os.Getenv("JWT_SECRET")
	jwtService := service.NewJWTService(secret, 24*time.Hour)

	// Middleware
	authMiddleware := middleware.NewAuthMiddleware(*jwtService)

	// Auth service & handler
	authService := service.NewAuthService(userRepo, roleRepo, jwtService)
	authHandler := handler.NewAuthHandle(authService)

	// Other handlers
	userHandle := handler.NewUserHandler(userService)
	roleHandle := handler.NewRoleHandle(roleService)

	// Setup router — pass authHandler too
	router := router.SetupRouter(userHandle, roleHandle, authHandler, authMiddleware.Handle())

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
