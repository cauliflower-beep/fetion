package service

import (
	"fetion/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"strconv"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} pong
// @Router /index [get]
func GetIndex(ctx *gin.Context) {
	index, err := template.ParseFiles("index.html", "views/chat/head.html")
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

// ToChat 跳转聊天主页面
func ToChat(ctx *gin.Context) {
	index, err := template.ParseFiles(
		"views/chat/index.html",
		"views/chat/head.html",
		"views/chat/tabmenu.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/main.html",
		"views/chat/createcom.html",
		"views/chat/userinfo.html",
		"views/chat/foot.html",
	)
	if err != nil {
		panic(err)
	}
	var user models.UserBasic
	uid, _ := strconv.Atoi(ctx.Query("userId"))
	token := ctx.Query("token")
	user.Uid = int64(uid)
	user.Identity = token
	// 登录之后，应该把uid和token传过来
	err = index.Execute(ctx.Writer, user)
	fmt.Println(err)
}
