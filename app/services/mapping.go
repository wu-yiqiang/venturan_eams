package services

import (
	"fmt"
	"venturan/app/common"
	"venturan/app/common/request"
	"venturan/app/models"
	"venturan/global"
)

type mappingService struct {
}

var MappingService = new(mappingService)

func (mappingService *mappingService) MappingPage(params request.CommonPageQueryForm) (err error, results *PaginationResult) {
	var mappings []models.Mapping
	query := global.App.DB.Scopes(QueryIsNoeDeleted).Where("name LIKE ? OR code LIKE ? OR description LIKE ?", "%"+params.Search+"%", "%"+params.Search+"%", "%"+params.Search+"%").Find(&mappings)
	rolePageResults, err := Pagination(query, params.PageNo, params.PageSize, &mappings)
	if err != nil {
		return
	}
	return nil, rolePageResults
}

func (mappingService *mappingService) MappingCreate(params request.MappingCreate) (err error) {
	mapping := models.Mapping{Name: models.Name{Name: params.Name}, Code: models.Code{Code: params.Code}, Value: params.Value, Color: params.Color, BackgroundColor: params.BackgroundColor, Description: models.Description{Description: params.Description}}
	result := global.App.DB.Create(&mapping)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}

func (mappingService *mappingService) MappingQueryDetails(mappingId string) (err error, mapping models.Mapping) {
	result := global.App.DB.Where("id = ?", mappingId).Find(&mapping)
	if result.Error != nil {
		return result.Error, models.Mapping{}
	}
	return nil, mapping
}

func (mappingService *mappingService) MappingUpdate(params request.MappingUpdate) (err error) {
	mapping := models.Mapping{Name: models.Name{Name: params.Name}, Code: models.Code{Code: params.Code}, Description: models.Description{Description: params.Description}, Color: params.Color, BackgroundColor: params.BackgroundColor}
	result := global.App.DB.Where("id = ?", params.ID).Updates(&mapping)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (mappingService *mappingService) MappingDelete(mappingId string) (err error) {
	result := global.App.DB.Model(&models.Mapping{}).Where("id = ?", mappingId).Update("is_deleted", common.Deleted)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}
