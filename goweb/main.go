package main

import (
	"firstproject/goweb/Internal/controller"
	"firstproject/goweb/global"
	"firstproject/goweb/inxit"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.POST("/login", controller.Login)
	inxit.Iniit()
	fmt.Println("-----------------")
	global.DB = inxit.DB
	port := ":8080"
	log.Println("端口号", port)
	err := r.Run(port)
	if err != nil {
		return
	}
}
