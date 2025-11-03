package services

import (
	"errors"
	"fmt"
	"time"
	"venturan/app/models"
	"venturan/global"
	"venturan/global/serviceErrors"
)

type parkService struct {
}

var ParkService = new(parkService)

func (parkService *parkService) ParkCountService() (data map[string]uint16, err error) {
	var datas = make(map[string]uint16)
	var used int64
	err = global.App.DB.Where("in_time is not null and out_time is null").Find(&models.ParkRecord{}).Count(&used).Error
	datas["used"] = uint16(used)
	datas["sum"] = global.App.Config.App.ParkingSpaceCapacity
	datas["remain"] = datas["sum"] - uint16(used)
	if err != nil {
		err = errors.New(serviceErrors.QueryError.Msg)
		return datas, err
	}
	return datas, nil
}

func (parkService *parkService) ParkEnterService(PlateNumberId uint) (parkRecord models.ParkRecord, err error) {
	parkRecord = models.ParkRecord{PlateNumberId: PlateNumberId, InTime: time.Now()}
	err = global.App.DB.Create(&parkRecord).Error
	fmt.Println("sss", err)
	if err != nil {
		err = errors.New(serviceErrors.CreateError.Msg)
		return
	}
	return
}

func (parkService *parkService) ParkExitService(PlateNumberId uint) (err error) {
	err = global.App.DB.Where("plate_number_id = ? and out_time is null and in_time is not null", PlateNumberId).Find(&models.ParkRecord{}).Update("out_time", time.Now()).Error
	if err != nil {
		err = errors.New(serviceErrors.UpdateError.Msg)
		return
	}
	return
}
