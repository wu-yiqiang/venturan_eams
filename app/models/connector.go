package models

type Connector struct {
	ID
	Name
	Method uint   `json:"method"`
	Path   string `json:"path" Gorm:"not null;comment:路径"`
	Description
	Timestamps
	IsDeleted
}
