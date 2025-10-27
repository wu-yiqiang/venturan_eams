package models

import (
	"time"
)

// 自增ID主键
type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

type Name struct {
	Name string `json:"name" gorm:"not null;comment:字段名称"`
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
	// 0 删除 1 使用中 2 归档中
	IsDeleted int8 `json:"is_deleted" gorm:"default:1;index"`
}

type Description struct {
	Description string `json:"description" gorm:"type:varchar(2500)"`
}

type Image struct {
	Image string `json:"image"`
}
