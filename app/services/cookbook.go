package services

import (
	"errors"
	"time"
	"venturan/app/models"
	"venturan/global"
	"venturan/global/serviceErrors"
)

type cookbookService struct {
}

var CookbookService = new(cookbookService)

// Register 注册
func (CookbookService *cookbookService) CookbookService() (err error, menus []models.CookBook) {
	err = global.App.DB.Where("created_at < ?", time.Now()).Find(&menus).Error
	if err != nil {
		err = errors.New(serviceErrors.QueryError.Msg)
		return
	}
	return
}
