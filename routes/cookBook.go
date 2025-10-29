package routes

import (
	"github.com/gin-gonic/gin"
	"venturan/app/controllers"
	"venturan/app/middleware"
	"venturan/app/services"
)

// Cook Book路由
func SetCookBookGroupRoutes(router *gin.RouterGroup) {
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/menus", controllers.CookBooks)
	}
}
