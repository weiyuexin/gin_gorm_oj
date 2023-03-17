package router

import "github.com/gin-gonic/gin"
import (
	docs "gin_gorm_oj/docs"
	"gin_gorm_oj/service"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	//Swagger配置
	//docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//路由规则
	r.GET("/ping", service.Ping)
	r.GET("/problem-list", service.GetProblemList)

	return r
}
