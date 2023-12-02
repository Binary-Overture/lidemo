package main

import (
	"Todolist/config"
	"Todolist/repository/db/dao"
	"Todolist/router"
)

func main() {
	loading()
	//创建一个路由
	r := router.NewRouter()
	//run路由
	_ = r.Run(config.Config.System.HttpPort)
}

func loading() {
	//初始化yaml文件
	config.InitConfig()
	//初始化mysql
	dao.MysqlInit()
}
