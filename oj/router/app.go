package router

import (
	_ "cncyx.xyz/docs"
	"cncyx.xyz/middlewares"
	"cncyx.xyz/service"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())
	v1 := r.Group("v1")
	{
		//swagger 配置
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		////问题
		//r.GET("/problemTest", service.GetProblemList)
		r.Any("/problem-list", service.GetProblemDetail)
		////用户
		r.Any("/user-detail", service.GetUserDetail) //用了Any之后好了
		r.POST("/login", service.Login)              //√
		r.POST("/send-code", service.SendEmailCode)  //x
		//提交
		r.GET("/problem-submit", service.GetSubmitList) //x
	}
	fmt.Println(v1)
	return r
}
