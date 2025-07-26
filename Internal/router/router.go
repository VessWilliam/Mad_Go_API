package router

import (
	"rest_api_gin/internal/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(userHandle *handler.UserHandle,
	roleHandle *handler.RoleHandle) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiGroup := router.Group("/api")

	//User Route
	apiGroup.GET("/getall", userHandle.GetUser)
	apiGroup.POST("/register", userHandle.RegisterUser)
	apiGroup.GET("/getbyid/:id", userHandle.GetById)

	//Role Route
	apiGroup.POST("/register_role", roleHandle.RegisterRole)

	return router
}
