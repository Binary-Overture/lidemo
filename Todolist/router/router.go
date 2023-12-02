package router

import (
	"Todolist/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	//gin使用中间件
	ginRouter.Use(middleware.Cors())
	//在gin框架中创建一个新的存储方法
	store := cookie.NewStore([]byte("something-very-secret"))
	//将store加入到gin中，这样就可以在处理HTTP时使用Session了
	ginRouter.Use(sessions.Sessions("MySession", store))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
	}
	return ginRouter
}
