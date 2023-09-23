package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	UUID     int64
	Username string
	Password string
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if isValid(username, password) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	} else {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "failed",
		})
	}
}

func isValid(username string, password string) bool {
	var users, test User
	DB.Model(&User{}).First(&users, "username=?", username, "password=?", password)
	if test == users {
		log.Fatal("没有数据，账号或密码错误")
		return false
	}
	return true
}
