package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func login(c *gin.Context) {

}

func main() {
	r := gin.Default()
	r.GET("/login", login)

	port := ":8080"
	log.Println("端口号", port)
	err := r.Run(port)
	if err != nil {
		return
	}
}
