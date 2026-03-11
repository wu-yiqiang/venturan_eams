package models

import (
	"time"
)

// 自增ID主键
type ID struct {
	ID uint `json:"id" Gorm:"primaryKey"`
}

type Name struct {
	Name string `json:"name" Gorm:"not null;comment:名称"`
}

type Code struct {
	Code string `json:"code" Gorm:"not null;comment:唯一编码"`
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
	IsDeleted int8 `json:"is_deleted" gorm:"default:0;index"`
}

type Description struct {
	Description string `json:"description" gorm:"type:varchar(2500)"`
}

type SerialNumber struct {
	SerialNumber string `json:"serialNumber" gorm:"type:varchar(20)"`
}

type Status struct {
	Status int8 `json:"status" gorm:"index"`
}

type File struct {
	FileName         string `json:"fileName" Gorm:"not null;comment:文件名"`
	FilePath         string `json:"filePath" Gorm:"not null;comment:文件路径"`
	OriginalFileName string `json:"originalFileName" Gorm:"not null;comment:源文件名"`
}

type Image struct {
	Image string `json:"image"`
}
