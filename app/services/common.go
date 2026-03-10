package services

import (
	"venturan/app/common"

	"gorm.io/gorm"
)

type PaginationResult struct {
	Data     interface{} `json:"data"`
	Total    int64       `json:"total"`
	PageNo   int         `json:"pageNo"`
	PageSize int         `json:"pageSize"`
}

func Pagination[T any](db *gorm.DB, pageNo, pageSize int, result *[]T) (*PaginationResult, error) {
	var total int64
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	db.Model(result).Count(&total)
	offset := (pageNo - 1) * pageSize
	err := db.Offset(offset).Limit(pageSize).Find(result).Error
	if err != nil {
		return nil, err
	}
	return &PaginationResult{
		Data:     result,
		Total:    total,
		PageNo:   pageNo,
		PageSize: pageSize,
	}, nil
}

func QueryIsNoeDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("is_deleted = ?", common.NotDeleted)
}
