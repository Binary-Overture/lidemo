package dao

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID     int64  `gorm:"column:uuid"` // 主键
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	//创建时间，时间戳

}
