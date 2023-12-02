package dao

import "gorm.io/gorm"

type TaskDao struct {
	*gorm.DB
}
