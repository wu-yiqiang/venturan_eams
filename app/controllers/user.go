package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"venturan/app/common/request"
	"venturan/app/common/response"
	"venturan/app/services"
	"venturan/global"
	"venturan/global/serviceErrors"
	"venturan/utils"
)

// @Summary 用户注册
// @Description 用户注册
// @Tags 注册接口
// @ID /user/register
// @Accept  json
// @Produce  json
// @Param body body request.Register true "body"
// @Success 200 {object} response.Response{data=request.Register} "success"
// @Router /user/register [post]
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

// @Summary 用户登陆
// @Description 用户登陆
// @Tags 登录接口
// @ID /user/login
// @Accept  json
// @Produce  json
// @Param body body request.Login true "body"    # [值得名称] body [值得类型] [是否必传] "[返回值名称]"
// @Success 200 {object} response.Response{data=request.Login} "success"
// @Router /user/login [post]
func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}
	if err, user := services.UserService.Login(form); err != nil {
		response.Fail(c, serviceErrors.UserIsNotExistOrPasswordError)
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.Fail(c, serviceErrors.TokentokenIssuanceFailed)
			return
		}
		userInfo := map[string]interface{}{}
		userInfoJson, _ := json.Marshal(user)
		json.Unmarshal(userInfoJson, &userInfo)
		userInfo["token"] = tokenData
		error := utils.SetRedisCache(tokenData, userInfo, time.Duration(global.App.Config.Jwt.JwtTtl)*time.Second)
		if error != nil {
			fmt.Println("存入缓存失败")
		}
		response.Success(c, userInfo)
	}
}

func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		// response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}

func Logout(c *gin.Context) {
	//err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	//if err != nil {
	//	response.BusinessFail(c, "登出失败")
	//	return
	//}
	//response.Success(c, nil)
}
