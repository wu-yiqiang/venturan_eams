package models

import (
	"gorm.io/gorm"
	"time"
)

// 自增ID主键
type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

// 创建、更新时间
type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateBy struct {
	CreatedBy time.Time `json:"created_by"`
	UpdatedBy time.Time `json:"updated_by"`
}

// 软删除
type IsDeleted struct {
	IsDeleted gorm.DeletedAt `json:"is_deleted" gorm:"index"`
}
