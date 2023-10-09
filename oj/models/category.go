package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	Name     string `gorm:"column:name;type:varchar(100);" json:"name"`
	ParentId string `gorm:"column:parent_id;type:int;" json:"parent_id"`
}

func (Category *Category) TableName() string {
	return "category"
}
