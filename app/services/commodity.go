package services

import (
	"fmt"
	"venturan/app/common"
	"venturan/app/common/request"
	"venturan/app/models"
	"venturan/global"
)

type commodityService struct {
}

var CommodityService = new(commodityService)

func (commodityService *commodityService) CommodityPage(params request.CommonPageQueryForm) (err error, results *PaginationResult) {
	var commodities []models.Commodity
	query := global.App.DB.Scopes(QueryIsNoeDeleted).Where("name LIKE ? OR code LIKE ? OR description LIKE ?", "%"+params.Search+"%", "%"+params.Search+"%", "%"+params.Search+"%").Find(&commodities)
	rolePageResults, err := Pagination(query, params.PageNo, params.PageSize, &commodities)
	if err != nil {
		return
	}
	return nil, rolePageResults
}

func (commodityService *commodityService) CommodityCreate(params request.CommodityCreate) (err error) {
	commodity := models.Commodity{Name: models.Name{Name: params.Name}, Code: models.Code{Code: params.Code}, Price: params.Price, Sales: params.Sales, Inventory: params.Inventory, FileName: models.FileName{params.FileName}, Status: models.Status{params.Status}, Description: models.Description{Description: params.Description}}
	result := global.App.DB.Create(&commodity)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}

func (commodityService *commodityService) CommodityQueryDetails(commodityId string) (err error, commodity models.Commodity) {
	result := global.App.DB.Where("id = ?", commodityId).Find(&commodity)
	if result.Error != nil {
		return result.Error, models.Commodity{}
	}
	return nil, commodity
}

func (commodityService *commodityService) CommodityUpdate(params request.CommodityUpdate) (err error) {
	commodity := models.Commodity{Name: models.Name{Name: params.Name}, Code: models.Code{Code: params.Code}, Price: params.Price, Sales: params.Sales, Inventory: params.Inventory, FileName: models.FileName{params.FileName}, Status: models.Status{params.Status}, Description: models.Description{Description: params.Description}}
	result := global.App.DB.Where("id = ?", params.ID).Updates(&commodity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (commodityService *commodityService) CommodityDelete(commodityId string) (err error) {
	result := global.App.DB.Model(&models.Commodity{}).Where("id = ?", commodityId).Update("is_deleted", common.Deleted)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}
