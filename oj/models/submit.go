package models

import "gorm.io/gorm"

type Submit struct {
	gorm.Model
	Identity        string `gorm:"column:identity;type:varchar(36);" json:"identity"`
	ProblemIdentity string `gorm:"column:problem_identity;type:varchar(36);" json:"problem_identity"`
	UserIdentity    string `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`
	CodePath        string `gorm:"column:code_path;type:varchar(255);" json:"code_path"`
	Status          int8   `gorm:"column:status;type:tinyint" json:"status"`
}

func (submit *Submit) TableName() string {
	return "submit"
}
