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
	// swagger
	docs.SwaggerInfo.BasePath = "" // 设置基础路由 比如这里设置为空 后面访问接口就是从空开始 "/index"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	/*
		静态资源 参数一是api 参数二是文件夹路径
		其中参数二需要注意：为来自给定文件系统`根`的文件提供服务,所以这里 asset 位于根目录，直接写即可
		也就是为静态文件做映射，通过项目url地址 /asset 访问 asset/下的静态文件
	*/
	r.Static("/asset", "asset/")
	/*
		LoadHTMLGlob 这个方法只能使用一次 多次调用的话 只有最后一次调用生效
		views/**/ /* 似乎不读取views里的文件 从两篇博客中看到的 亲自试下views/index.html 也没读到.
	可以试下读取模板文件的另外一种方法：
	r.LoadHTMLFiles(
		"views/index.html",
		"views/index2.html"
	)
	*/
	//r.LoadHTMLFiles("index.html")
	r.LoadHTMLGlob("views/**/*")
	// 首页
	r.GET("/index", service.GetIndex)
	r.GET("/register", service.Register)
	r.GET("/toChat", service.ToChat)
	// 用户相关
	r.POST("/user/getUsers", service.GetUsers)
	r.POST("/user/getUserByNameAndPwd", service.GetUserByNameAndPwd)
	r.POST("/user/createUser", service.CreateUser) // 注册
	r.POST("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)
	// 消息相关
	r.GET("msg/sendMsg", service.SendMsg)               // Get 因为页面上只对应一个按钮
	r.GET("msg/sendPrivateMsg", service.SendPrivateMsg) // Get 因为页面上只对应一个按钮
	return r
}
