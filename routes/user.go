package routes

import (
	"venturan/app/controllers"
	"venturan/app/middleware"
	"venturan/app/services"

	"github.com/gin-gonic/gin"
)

// user
func SetUserGroupRoutes(router *gin.RouterGroup) {
	router.POST("/register", controllers.UserRegister)
	router.POST("/login", controllers.UserLogin)
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.DELETE("/:user_id", controllers.UserDelete)
		authRouter.POST("/info", controllers.UserInfo)
		authRouter.POST("/logout", controllers.UserLogout)
		authRouter.POST("/page", controllers.UserPage)
		authRouter.POST("/lists", controllers.UserList)
	}
}
