package models

import "gorm.io/gorm"

type SubmitBasic struct {
	gorm.Model
	Identity        string        `gorm:"column:identity;type:varchar(36);" json:"identity"`
	ProblemIdentity string        `gorm:"column:problem_identity;type:varchar(36);" json:"problem_identity"`
	ProblemBasic    *ProblemBasic `gorm:"foreign:identity;reference:problem_identity;"` //关联问题基础表
	UserIdentity    string        `gorm:"column:user_identity;type:varchar(36);" json:"user_identity"`
	UserBasic       *UserBasic    `gorm:"column:identity;reference:user_identity"`
	CodePath        string        `gorm:"column:code_path;type:varchar(255);" json:"code_path"`
	Status          int8          `gorm:"column:status;type:tinyint" json:"status"`
}

func (submit *SubmitBasic) TableName() string {
	return "submit_basic"
}

func GetSubmitListByAll(problemIdentity string, userIdentity string, status int) *gorm.DB {
	tx := DB.Model(new(SubmitBasic)).Preload("ProblemBasic").Preload("UserBasic")

	if problemIdentity != "" {
		tx.Where("problem_identity = ?", problemIdentity)
	}
	if userIdentity != "" {
		tx.Where("user_identity = ?", userIdentity)
	}
	if status != -1 {
		tx.Where("status = ?", status)
	}
	return tx
}
