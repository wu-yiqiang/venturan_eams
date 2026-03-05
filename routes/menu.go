package routes

import (
	"venturan/app/controllers"
	"venturan/app/middleware"
	"venturan/app/services"

	"github.com/gin-gonic/gin"
)

// user
func SetMenuGroupRoutes(router *gin.RouterGroup) {
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/page", controllers.MenuPages)
		authRouter.POST("/list", controllers.MenuLists)
		authRouter.DELETE("/:menu_id", controllers.MenuDelete)
	}
}
