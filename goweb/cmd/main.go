package main

import (
	"firstproject/goweb/inxit"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

var (
	DB *gorm.DB
)

func main() {
	r := gin.Default()
	r.GET("/login", login)
	inxit.Iniit()
	fmt.Println("-----------------")
	DB = inxit.DB
	port := ":8080"
	log.Println("端口号", port)
	err := r.Run(port)
	if err != nil {
		return
	}
}
