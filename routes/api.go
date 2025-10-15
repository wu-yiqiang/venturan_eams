package routes

import (
	"github.com/gin-gonic/gin"
	"venturan/app/controllers/app"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/user/register", app.Register)
}
