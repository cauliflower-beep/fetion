package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} pong
// @Router /index [get]
func GetIndex(ctx *gin.Context) {
	index, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	err = index.Execute(ctx.Writer, "index")
	fmt.Println(err)
}

// Register
// @Tags 注册页面
// @Success 200 {string} pong
// @Router /register [get]
func Register(ctx *gin.Context) {
	index, err := template.ParseFiles("views/user/register.html")
	if err != nil {
		panic(err)
	}
	err = index.Execute(ctx.Writer, "register")
	fmt.Println(err)
}
