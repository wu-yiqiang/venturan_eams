package routes

import (
	"venturan/app/controllers"
	"venturan/app/middleware"
	"venturan/app/services"

	"github.com/gin-gonic/gin"
)

// user
func SetMappingGroupRoutes(router *gin.RouterGroup) {
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/page", controllers.MappingPage)
		authRouter.POST("/lists", controllers.MappingList)
		authRouter.POST("/create", controllers.MappingCreate)
		authRouter.GET("/details/:mapping_id", controllers.MappingDetails)
		authRouter.POST("/update", controllers.MappingUpdate)
		authRouter.DELETE("/delete/:mapping_id", controllers.MappingDelete)
	}
}
