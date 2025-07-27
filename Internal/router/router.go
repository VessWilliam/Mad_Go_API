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
	apiGroup.PUT("/update_user", userHandle.UpdateUser)
	apiGroup.GET("/getbyid_user/:id", userHandle.GetById)
	apiGroup.POST("/register_user", userHandle.RegisterUser)
	apiGroup.PUT("/assign-roles", userHandle.AssignRolesToUser)

	//Role Route
	apiGroup.GET("/get_roles", roleHandle.GetRoles)
	apiGroup.PUT("/update_role", roleHandle.UpdateRole)
	apiGroup.POST("/register_role", roleHandle.RegisterRole)
	apiGroup.GET("/getbyid_role/:id", roleHandle.GetRolesById)
	apiGroup.DELETE("/deletebyid_role/:id", roleHandle.DeleteById)

	return router
}
