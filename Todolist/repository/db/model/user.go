package model

import (
	"gorm.io/gorm"
)

// UserModel 用户模型
type UserModel struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string
}

func (*UserModel) TableName() string {
	return "user"
}
