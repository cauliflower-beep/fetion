package router

import (
	"fetion/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", service.GetIndex)
	return r
}
