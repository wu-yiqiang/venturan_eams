package routes

import (
	"venturan/app/controllers"
	"venturan/app/middleware"
	"venturan/app/services"

	"github.com/gin-gonic/gin"
)

// user
func SetRoleGroupRoutes(router *gin.RouterGroup) {
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/page", controllers.RolePage)
		authRouter.POST("/lists", controllers.RoleList)
	}
}
