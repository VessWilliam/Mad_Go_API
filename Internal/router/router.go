package router

import (
	"rest_api_gin/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandle *handler.UserHandle) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/register", userHandle.RegisterUser)
	}

	return router
}
