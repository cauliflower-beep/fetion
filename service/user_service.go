package service

import (
	"fetion/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUsers
// @Summary 获取用户列表
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUsers [get]
func GetUsers(ctx *gin.Context) {
	users := models.GetUserList()
	ctx.JSON(http.StatusOK, gin.H{
		"name": users,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param pwd  query string false "密码"
// @param rePwd query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(ctx *gin.Context) {
	// 解析URL参数
	user := &models.UserBasic{}
	name := ctx.DefaultQuery("name", "忧郁的小黄鹂")
	pwd := ctx.DefaultQuery("pwd", "test123")     // 密码
	rePwd := ctx.DefaultQuery("rePwd", "test123") // 确认密码
	if pwd != rePwd {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "密码不一致，请重新确认！",
		})
		return
	}
	user.Name = name
	user.Pwd = pwd
	// 创建用户
	models.CreateUser(user)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "用户创建成功",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(ctx *gin.Context) {
	// 解析URL参数
	user := &models.UserBasic{}
	id, _ := strconv.Atoi(ctx.Query("id"))
	user.ID = uint(id)

	// 删除用户
	models.DeleteUser(user)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "用户删除成功",
	})
}

// UpdateUser
// @Summary 更新用户资料
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "用户名"
// @param pwd  formData string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(ctx *gin.Context) {
	// 解析id
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	name := ctx.PostForm("name")
	pwd := ctx.PostForm("pwd")
	var user models.UserBasic
	user.ID = uint(id)
	user.Name = name
	user.Pwd = pwd
	// 更新用户数据
	models.UpdateUser(user)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "用户数据更新成功",
	})
}
