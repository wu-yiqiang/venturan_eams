package services

import (
	"errors"
	"venturan/app/models"
	"venturan/global"
	"venturan/global/serviceErrors"
)

type plateNumberService struct {
}

var PlateNumberService = new(plateNumberService)

// Register 注册
func (plateNumberService *plateNumberService) FindPlateNumber(plateNum string) (plateNumber models.PlateNumber, err error) {
	err = global.App.DB.Where("name = ?", plateNum).First(&plateNumber).Error
	if err != nil {
		err = errors.New(serviceErrors.QueryError.Msg)
		return
	}
	return
}
