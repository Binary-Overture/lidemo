package controller

import (
	"firstproject/goweb/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	UUID     int64
	Username string
	Password string
}

func Register(c *gin.Context) {
	//用户名,密码,性别,年龄
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("username:", username)
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
	global.DB.Where("username=? and password=?", username, password).First(&users)
	if test == users {
		fmt.Println("没有数据，账号或密码错误")
		return false
	}
	return true
}
