package router

import (
	"rest_api_gin/internal/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(
	userHandle *handler.UserHandle,
	roleHandle *handler.RoleHandle,
	authHandle *handler.AuthHandler,
	authMiddleware gin.HandlerFunc,
) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiGroup := router.Group("/api")

	// Protected routes
	JWTProtected := apiGroup.Group("/")
	JWTProtected.Use(authMiddleware)

	// Public routes
	apiGroup.POST("/login", authHandle.Login)
	apiGroup.GET("/get_roles", roleHandle.GetRoles)
	apiGroup.GET("/get_users", userHandle.GetUsers)
	apiGroup.POST("/register_user", userHandle.RegisterUser)
	apiGroup.POST("/register_role", roleHandle.RegisterRole)

	// User Routes
	JWTProtected.GET("/getbyid_user/:id", userHandle.GetById)
	JWTProtected.PUT("/update_user", userHandle.UpdateUser)
	JWTProtected.PUT("/assign-roles", userHandle.AssignRolesToUser)

	// Role Routes
	JWTProtected.GET("/getbyid_role/:id", roleHandle.GetRolesById)
	JWTProtected.PUT("/update_role", roleHandle.UpdateRole)
	JWTProtected.DELETE("/deletebyid_role/:id", roleHandle.DeleteById)

	return router
}
