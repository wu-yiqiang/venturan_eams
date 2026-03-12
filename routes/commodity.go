package routes

import (
	"venturan/app/controllers"
	"venturan/app/middleware"
	"venturan/app/services"

	"github.com/gin-gonic/gin"
)

// 商品
func SetCommodityGroupRoutes(router *gin.RouterGroup) {
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/page", controllers.CommodityPage)
		authRouter.POST("/create", controllers.CommodityCreate)
		authRouter.GET("/details/:commodity_id", controllers.CommodityDetails)
		authRouter.POST("/update", controllers.CommodityUpdate)
		authRouter.DELETE("/delete/:commodity_id", controllers.CommodityDelete)
	}
}
