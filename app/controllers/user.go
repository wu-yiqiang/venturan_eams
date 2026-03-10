package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"venturan/app/common/request"
	"venturan/app/common/response"
	"venturan/app/services"
	"venturan/global"
	"venturan/global/serviceErrors"
	"venturan/utils"

	"github.com/gin-gonic/gin"
)

// @Summary 用户注册
// @Description 用户注册
// @Tags 用户管理
// @ID /user/register
// @Accept  json
// @Produce  json
// @Param body body request.Register true "body"
// @Success 200 {object} response.Response{data=models.User} "success"
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
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
// @Tags 用户管理
// @ID /user/login
// @Accept  json
// @Produce  json
// @Param body body request.Login true "body"
// @Success 200 {object} response.Response{data=models.User} "success"
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
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

// @Summary 用户删除
// @Description 用户删除
// @Tags 用户管理
// @ID /user/{id}
// @Accept  json
// @Produce  json
// @Param id path int 1 "id"
// @Router /user/{id} [delete]
func UserDelete(c *gin.Context) {
	userId := c.Param("user_id")
	if userId == "" {
		response.Fail(c, serviceErrors.UserIdNotEmpty)
		return
	}
	id, err := strconv.Atoi(userId)
	if err != nil {
		response.Fail(c, serviceErrors.DeleteError)
		return
	}
	err, user := services.UserService.Delete(uint(id))
	if err != nil {
		response.Fail(c, serviceErrors.DeleteError)
		return
	}
	response.Success(c, user)
}

func UserInfo(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		// response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}

func UserLogout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	err := services.UserService.UserLogout(token)
	if err != nil {
		response.Fail(c, serviceErrors.UserExitFailed)
		return
	}
	response.Success(c, nil)
}

// @Summary 用户分页查询
// @Description 用户分页查询
// @Tags 用户管理
// @ID /user/page
// @Accept  json
// @Produce  json
// @Router /user/page [post]
// @Param body body request.UserPageQueryForm true "body"
// @Success 200 {object} response.Response{data=[]models.User} "success"
func UserPage(c *gin.Context) {
	var userPageQueryForm request.UserPageQueryForm
	//var paginationResult services.PaginationResult
	if err := c.ShouldBindJSON(&userPageQueryForm); err != nil {
		response.Fail(c, request.GetErrorMsg(userPageQueryForm, err))
		return
	}
	err, paginationResult := services.UserService.UserPage(userPageQueryForm)
	if err != nil {
		response.Fail(c, serviceErrors.UserIsNotExistOrPasswordError)
		return
	}
	response.Success(c, paginationResult)
}

func UserList(c *gin.Context) {}
