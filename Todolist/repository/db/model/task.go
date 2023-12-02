package model

import (
	"gorm.io/gorm"
)

// TaskModel 任务模型
type TaskModel struct {
	gorm.Model
	User      UserModel `gorm:"ForeignKey:Uid"`
	Uid       uint      `gorm:"not null"`
	Title     string    `gorm:"index;not null"`
	Status    int       `gorm:"default:0"`
	Content   string    `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64 `gorm:"default:0"`
}

func (*TaskModel) TableName() string {
	return "task"
}
