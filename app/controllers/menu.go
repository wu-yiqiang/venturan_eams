package controllers

import (
	"venturan/app/common/request"
	"venturan/app/common/response"

	"github.com/gin-gonic/gin"
)

// @Summary 菜单列表分页
// @Description 菜单列表分页
// @Tags 菜单管理
// @ID /menu/page
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=[]models.Menu} "success"
// @Router /menu/page [post]
func MenuPages(c *gin.Context) {
	var form request.MenuPage
	if err := c.ShouldBindJSON(&form); err != nil {
		response.Fail(c, request.GetErrorMsg(form, err))
		return
	}
	//if err, user := services.UserService.Register(form); err != nil {
	//	response.Fail(c, serviceErrors.UserCreateFailed)
	//} else {
	//	response.Success(c, user)
	//}
}

// @Summary 菜单列表
// @Description 菜单列表
// @Tags 菜单管理
// @ID /menu/list
// @Accept  json
// @Produce json
// @Success 200 {object} response.Response{data=[]models.Menu} "success"
// @Router /menu/list [post]
func MenuLists(c *gin.Context) {
	//plateNumber := c.Param("plateNumber")
	//plateInfo, findPlateIdError := FindPlateInfoByNumber(plateNumber)
	//if findPlateIdError != nil {
	//	response.Fail(c, serviceErrors.LicensePlateNotRegister)
	//	return
	//}
	//parkRecord, err := services.ParkService.ParkEnterService(plateInfo.ID.ID)
	//if err != nil {
	//	response.Fail(c, serviceErrors.CreateError)
	//	return
	//}
	//response.Success(c, parkRecord)

}

// @Summary 菜单删除
// @Description 菜单删除
// @Tags 菜单管理
// @ID /menu/{id}
// @Accept  json
// @Produce  json
// @Param id path int 1 "id"
// @Router /menu/{id} [delete]
func MenuDelete(c *gin.Context) {
	//menuId := c.Param("menu_id")
	//if menuId == "" {
	//	response.Fail(c, serviceErrors.UserIdNotEmpty)
	//	return
	//}
	//id, err := strconv.Atoi(menuId)
	//if err != nil {
	//	response.Fail(c, serviceErrors.DeleteError)
	//	return
	//}
	//err, user := services.UserService.Delete(uint(id))
	//if err != nil {
	//	response.Fail(c, serviceErrors.DeleteError)
	//	return
	//}
	//response.Success(c, user)
}
