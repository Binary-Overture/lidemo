package main

import (
	"firstproject/goweb/Internal/controller"
	"firstproject/goweb/Internal/dao"
	"firstproject/goweb/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.POST("/login", controller.Login)
	dao.Iniit()
	fmt.Println("-----------------")
	global.DB = dao.DB
	port := ":8080"
	log.Println("端口号", port)
	err := r.Run(port)
	if err != nil {
		return
	}
}
