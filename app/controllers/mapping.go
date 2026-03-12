package controllers

import (
	"venturan/app/common/request"
	"venturan/app/common/response"
	"venturan/app/services"
	"venturan/global/serviceErrors"

	"github.com/gin-gonic/gin"
)

// @Summary 字典分页查询
// @Description 字典分页查询
// @Tags 字典管理
// @ID /mapping/page
// @Accept  json
// @Produce  json
// @Router /mapping/page [post]
// @Param body body request.CommonPageQueryForm true "body"
// @Success 200 {object} response.Response{data=[]models.Mapping} "success"
func MappingPage(c *gin.Context) {
	var mappingPageQueryForm request.CommonPageQueryForm
	if err := c.ShouldBindJSON(&mappingPageQueryForm); err != nil {
		response.Fail(c, request.GetErrorMsg(mappingPageQueryForm, err))
		return
	}
	err, paginationResult := services.MappingService.MappingPage(mappingPageQueryForm)
	if err != nil {
		response.Fail(c, serviceErrors.QueryError)
		return
	}
	response.Success(c, paginationResult)
}

// @Summary 字典创建
// @Description 字典创建
// @Tags 字典管理
// @ID /mapping/create
// @Accept  json
// @Produce  json
// @Router /mapping/create [post]
// @Param body body request.MappingCreate true "body"
// @Success 200 {object} response.Response{} "success"
func MappingCreate(c *gin.Context) {
	var mappingForm request.MappingCreate
	if err := c.ShouldBindJSON(&mappingForm); err != nil {
		response.Fail(c, request.GetErrorMsg(mappingForm, err))
		return
	}
	err := services.MappingService.MappingCreate(mappingForm)
	if err != nil {
		response.Fail(c, serviceErrors.CreateError)
		return
	}
	response.Success(c, nil)
}

// @Summary 字典详情
// @Description 字典详情
// @Tags 字典管理
// @ID /mapping/details/{mappingId}
// @Accept  json
// @Produce  json
// @Router /mapping/details/{mappingId} [get]
// @Param mappingId path int 1 "mappingId"
// @Success 200 {object} response.Response{data=models.Mapping} "success"
func MappingDetails(c *gin.Context) {
	mappingId := c.Param("mapping_id")
	if mappingId == "" {
		response.Fail(c, serviceErrors.IdIsNotExist)
		return
	}
	err, mappingDetails := services.MappingService.MappingQueryDetails(mappingId)
	if err != nil {
		response.Fail(c, serviceErrors.QueryDetailsError)
		return
	}
	response.Success(c, mappingDetails)
}

// @Summary 字典更新
// @Description 字典更新
// @Tags 字典管理
// @ID /mapping/update
// @Accept  json
// @Produce  json
// @Router /mapping/update [post]
// @Param body body request.MappingCreate true "body"
// @Success 200 {object} response.Response{data=models.Mapping} "success"
func MappingUpdate(c *gin.Context) {
	var mappingForm request.MappingUpdate
	if err := c.ShouldBindJSON(&mappingForm); err != nil {
		response.Fail(c, request.GetErrorMsg(mappingForm, err))
		return
	}
	err := services.MappingService.MappingUpdate(mappingForm)
	if err != nil {
		response.Fail(c, serviceErrors.CreateError)
		return
	}
	response.Success(c, nil)
}

// @Summary 字典删除
// @Description 字典删除
// @Tags 字典管理
// @ID /mapping/delete/{mappingId}
// @Accept  json
// @Produce  json
// @Router /mapping/delete/{mappingId} [delete]
// @Param mappingId path int 1 "mappingId"
// @Success 200 {object} response.Response{} "success"
func MappingDelete(c *gin.Context) {
	mappingId := c.Param("mapping_id")
	if mappingId == "" {
		response.Fail(c, serviceErrors.IdIsNotEmpty)
		return
	}
	err := services.MappingService.MappingDelete(mappingId)
	if err != nil {
		response.Fail(c, serviceErrors.DeleteError)
		return
	}
	response.Success(c, nil)
}

// @Summary 字典类型查询
// @Description 字典类型查询
// @Tags 字典管理
// @ID /mapping/types
// @Accept  json
// @Produce  json
// @Router /mapping/types [post]
// @Param body body request.MappingType true "body"
// @Success 200 {object} response.Response{data=[]models.Mapping} "success"
func MappingTypes(c *gin.Context) {
	var mappingForm request.MappingType
	if err := c.ShouldBindJSON(&mappingForm); err != nil {
		response.Fail(c, request.GetErrorMsg(mappingForm, err))
		return
	}
	err, mappings := services.MappingService.MappingTypeList(mappingForm)
	if err != nil {
		response.Fail(c, serviceErrors.QueryError)
		return
	}
	response.Success(c, mappings)
}

func MappingList(c *gin.Context) {}
