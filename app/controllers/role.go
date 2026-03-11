package controllers

import (
	"venturan/app/common/request"
	"venturan/app/common/response"
	"venturan/app/services"
	"venturan/global/serviceErrors"

	"github.com/gin-gonic/gin"
)

// @Summary 角色分页查询
// @Description 角色分页查询
// @Tags 角色管理
// @ID /role/page
// @Accept  json
// @Produce  json
// @Router /role/page [post]
// @Param body body request.CommonPageQueryForm true "body"
// @Success 200 {object} response.Response{data=[]models.Role} "success"
func RolePage(c *gin.Context) {
	var rolePageQueryForm request.CommonPageQueryForm
	if err := c.ShouldBindJSON(&rolePageQueryForm); err != nil {
		response.Fail(c, request.GetErrorMsg(rolePageQueryForm, err))
		return
	}
	err, paginationResult := services.RoleService.RolePage(rolePageQueryForm)
	if err != nil {
		response.Fail(c, serviceErrors.QueryError)
		return
	}
	response.Success(c, paginationResult)
}

func RoleList(c *gin.Context) {}
