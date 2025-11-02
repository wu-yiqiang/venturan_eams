package routes

import (
	"github.com/gin-gonic/gin"
	"venturan/app/controllers"
	"venturan/app/middleware"
	"venturan/app/services"
)

// Cook Book路由
func SetParkGroupRoutes(router *gin.RouterGroup) {
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/count", controllers.ParkDetails)
		authRouter.POST("/enter/:plateNumber", controllers.ParkEnter)
		authRouter.POST("/exit/:plateNumber", controllers.ParkExit)
	}
}
