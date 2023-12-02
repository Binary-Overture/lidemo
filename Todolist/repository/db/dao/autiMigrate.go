package dao

import (
	"Todolist/repository/db/model"
)

func autoMigrate() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.TaskModel{}, &model.UserModel{})

	if err != nil {
		panic(err)
	}
}
