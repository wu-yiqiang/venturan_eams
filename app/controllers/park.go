package controllers

import (
	"github.com/gin-gonic/gin"
	"venturan/app/common/response"
	"venturan/app/models"
	"venturan/app/services"
	"venturan/global/serviceErrors"
)

// @Summary 车位详情
// @Description 车位详情
// @Tags 热门服务
// @ID /park/count
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response{data=models.CookBook} "success"
// @Router /park/count [post]
func ParkDetails(c *gin.Context) {
	parkCount, err := services.ParkService.ParkCountService()
	if err != nil {
		response.Fail(c, serviceErrors.QueryError)
		return
	}
	response.Success(c, parkCount)
}

// @Summary 入园记录
// @Description 入园记录
// @Tags 热门服务
// @ID /park/enter
// @Accept  json
// @Produce  json
// @Param plateNumber path string 1 "plateNumber"
// @Success 200 {object} response.Response{data=models.ParkRecord} "success"
// @Router /park/enter/{plateNumber} [post]
func ParkEnter(c *gin.Context) {
	plateNumber := c.Param("plateNumber")
	plateInfo, findPlateIdError := FindPlateInfoByNumber(plateNumber)
	if findPlateIdError != nil {
		response.Fail(c, serviceErrors.LicensePlateNotRegister)
		return
	}
	parkRecord, err := services.ParkService.ParkEnterService(plateInfo.ID.ID)
	if err != nil {
		response.Fail(c, serviceErrors.CreateError)
		return
	}
	response.Success(c, parkRecord)

}

// @Summary 出园记录
// @Description 出园记录
// @Tags 热门服务
// @ID /park/exit
// @Accept  json
// @Produce  json
// @Param plateNumber path string 1 "plateNumber"
// @Success 200 {object} response.Response{data=models.ParkRecord} "success"
// @Router /park/exit/{plateNumber} [post]
func ParkExit(c *gin.Context) {
	plateNumber := c.Param("plateNumber")
	plateInfo, findPlateIdError := FindPlateInfoByNumber(plateNumber)
	if findPlateIdError != nil {
		response.Fail(c, serviceErrors.LicensePlateNotRegister)
		return
	}
	err := services.ParkService.ParkExitService(plateInfo.ID.ID)
	if err != nil {
		response.Fail(c, serviceErrors.CreateError)
		return
	}
	response.Success(c, nil)

}

func FindPlateInfoByNumber(plateNum string) (plateInfo models.PlateNumber, err error) {
	plateInfo, err = services.PlateNumberService.FindPlateNumber(plateNum)
	if err != nil {
		return
	}
	return
}
