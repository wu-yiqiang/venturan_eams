package app

import (
	"github.com/gin-gonic/gin"
	"venturan/app/common/request"
	"venturan/app/common/response"
	"venturan/app/services"
	"venturan/global/serviceErrors"
)

// Register 用户注册
func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}
	if err, user := services.UserService.Register(form); err != nil {
		response.Fail(c, serviceErrors.UserCreateFailed)
	} else {
		response.Success(c, user)
	}
}
