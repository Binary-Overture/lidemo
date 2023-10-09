package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Identity   string `gorm:"column:identity;type:varchar(36);" json:"identity"`        //问题类的唯一标识
	CategoryId string `gorm:"column:category_id;type:varchar(255);" json:"category_id"` //分类id，已逗号分隔
	Title      string `gorm:"column:category;type:varchar(255);" json:"title"`          //文章标题
	Content    string `gorm:"column:content;type:text;" json:"content"`                 //文章正文
	MaxMem     int    `gorm:"column:max_mem;type:int;" json:"max_mem"`                  //最大的运行内存
	MaxRuntime int    `gorm:"column:max_runtime;type:int;" json:"max_runtime"`          //最大的运行时间
}

func (table *Problem) TableName() string {
	return "problem"
}