package routes

import (
	"github.com/gin-gonic/gin"
	"venturan/app/controllers"
	"venturan/app/middleware"
	"venturan/app/services"
)

// user
func SetUserGroupRoutes(router *gin.RouterGroup) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.DELETE("/:user_id", controllers.Delete)
		authRouter.POST("/info", controllers.Info)
		authRouter.POST("/logout", controllers.Logout)
	}
}
