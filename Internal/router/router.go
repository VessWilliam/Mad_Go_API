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
	apiGroup.GET("/get_users", userHandle.GetUsers)
	apiGroup.POST("/register_user", userHandle.RegisterUser)
	apiGroup.GET("/getbyid_user/:id", userHandle.GetById)

	//Role Route
	apiGroup.GET("/get_roles", roleHandle.GetRoles)
	apiGroup.POST("/register_role", roleHandle.RegisterRole)
	apiGroup.GET("/getbyid_role/:id", roleHandle.GetRolesById)
	apiGroup.DELETE("/deletebyid_role/:id", roleHandle.DeleteById)

	return router
}
