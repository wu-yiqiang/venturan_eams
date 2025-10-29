package controllers

import (
	"github.com/gin-gonic/gin"
	"venturan/app/common/response"
	"venturan/app/services"
	"venturan/global/serviceErrors"
)

// @Summary 热门服务
// @Description 热门服务
// @Tags 今日菜单
// @ID /cookbook/menus
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=models.CookBook} "success"
// @Router /cookbook/menus [post]
func CookBooks(c *gin.Context) {
	if err, cookBooks := services.CookbookService.CookbookService(); err != nil {
		response.Fail(c, serviceErrors.QueryError)
	} else {
		response.Success(c, cookBooks)
	}
}
