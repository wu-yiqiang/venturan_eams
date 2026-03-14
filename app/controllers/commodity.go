package controllers

import (
	"venturan/app/common/request"
	"venturan/app/common/response"
	"venturan/app/services"
	"venturan/global/serviceErrors"

	"github.com/gin-gonic/gin"
)

// @Summary 商品分页查询
// @Description 商品分页查询
// @Tags 商品管理
// @ID /commodity/page
// @Accept  json
// @Produce  json
// @Router /commodity/page [post]
// @Param body body request.CommonPageQueryForm true "body"
// @Success 200 {object} response.Response{data=[]models.Commodity} "success"
func CommodityPage(c *gin.Context) {
	var commodityPageQueryForm request.CommonPageQueryForm
	if err := c.ShouldBindJSON(&commodityPageQueryForm); err != nil {
		response.Fail(c, request.GetErrorMsg(commodityPageQueryForm, err))
		return
	}
	err, paginationResult := services.CommodityService.CommodityPage(commodityPageQueryForm)
	if err != nil {
		response.Fail(c, serviceErrors.QueryError)
		return
	}
	response.Success(c, paginationResult)
}

// @Summary 添加商品
// @Description 添加商品
// @Tags 商品管理
// @ID /commodity/create
// @Accept  json
// @Produce  json
// @Router /commodity/create [post]
// @Param body body request.MappingCreate true "body"
// @Success 200 {object} response.Response{} "success"
func CommodityCreate(c *gin.Context) {
	var commodityForm request.CommodityCreate
	if err := c.ShouldBindJSON(&commodityForm); err != nil {
		response.Fail(c, request.GetErrorMsg(commodityForm, err))
		return
	}
	err := services.CommodityService.CommodityCreate(commodityForm)
	if err != nil {
		response.Fail(c, serviceErrors.CreateError)
		return
	}
	response.Success(c, nil)
}

// @Summary 商品详情
// @Description 商品详情
// @Tags 商品管理
// @ID /commodity/details/{commodityId}
// @Accept  json
// @Produce  json
// @Router /commodity/details/{commodityId} [get]
// @Param commodityId path int 1 "commodityId"
// @Success 200 {object} response.Response{data=models.Commodity} "success"
func CommodityDetails(c *gin.Context) {
	commodityId := c.Param("commodity_id")
	if commodityId == "" {
		response.Fail(c, serviceErrors.IdIsNotExist)
		return
	}
	err, commodityDetails := services.CommodityService.CommodityQueryDetails(commodityId)
	if err != nil {
		response.Fail(c, serviceErrors.QueryDetailsError)
		return
	}
	response.Success(c, commodityDetails)
}

// @Summary 更新商品
// @Description 更新商品
// @Tags 商品管理
// @ID /commodity/update
// @Accept  json
// @Produce  json
// @Router /commodity/update [post]
// @Param body body request.MappingCreate true "body"
// @Success 200 {object} response.Response{data=models.Commodity} "success"
func CommodityUpdate(c *gin.Context) {
	var commodityForm request.CommodityUpdate
	if err := c.ShouldBindJSON(&commodityForm); err != nil {
		response.Fail(c, request.GetErrorMsg(commodityForm, err))
		return
	}
	err := services.CommodityService.CommodityUpdate(commodityForm)
	if err != nil {
		response.Fail(c, serviceErrors.CreateError)
		return
	}
	response.Success(c, nil)
}

// @Summary 商品上架
// @Description 商品上架
// @Tags 商品管理
// @ID /commodity/up
// @Accept  json
// @Produce  json
// @Router /commodity/up [post]
// @Param body body request.MappingCreate true "body"
// @Success 200 {object} response.Response{data=models.Commodity} "success"
func CommodityUpdateStatus(c *gin.Context) {
	var commodityForm request.CommodityUp
	if err := c.ShouldBindJSON(&commodityForm); err != nil {
		response.Fail(c, request.GetErrorMsg(commodityForm, err))
		return
	}
	err := services.CommodityService.CommodityUp(commodityForm)
	if err != nil {
		response.Fail(c, serviceErrors.UpdateError)
		return
	}
	response.Success(c, nil)
}

// @Summary 删除商品
// @Description 删除商品
// @Tags 商品管理
// @ID /commodity/delete/{commodityId}
// @Accept  json
// @Produce  json
// @Router /commodity/delete/{commodityId} [delete]
// @Param commodityId path int 1 "commodityId"
// @Success 200 {object} response.Response{} "success"
func CommodityDelete(c *gin.Context) {
	commodityId := c.Param("commodity_id")
	if commodityId == "" {
		response.Fail(c, serviceErrors.IdIsNotEmpty)
		return
	}
	err := services.CommodityService.CommodityDelete(commodityId)
	if err != nil {
		response.Fail(c, serviceErrors.DeleteError)
		return
	}
	response.Success(c, nil)
}
