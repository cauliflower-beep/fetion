package router

import (
	docs "fetion/docs"
	"fetion/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "" // 设置基础路由 比如这里设置为空 后面访问接口就是从空开始 "/index"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index", service.GetIndex)
	// 用户相关
	r.GET("/user/getUsers", service.GetUsers)
	r.GET("/user/getUserByNameAndPwd", service.GetUserByNameAndPwd)
	r.GET("/user/createUser", service.CreateUser)
	r.GET("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)
	return r
}
