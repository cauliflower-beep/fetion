package service

import (
	"fetion/models"
	"fetion/utils"
	"fmt"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
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
		"code": 0,
		"data": users,
	})
}

// GetUserByNameAndPwd
// @Summary 根据用户名和密码获取用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param pwd  query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserByNameAndPwd [post]
func GetUserByNameAndPwd(ctx *gin.Context) {
	name := ctx.Query("name")
	pwd := ctx.Query("pwd")
	user, msg := models.FindUserByNameAndPwd(name, pwd)
	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1, // 0-成功 -1-失败
			"msg":  msg,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
		"data": user,
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
	name := ctx.Query("name")
	// 不允许出现重复的用户名 后续就支持通过用户名登陆
	rec := models.FindUserByName(name)
	if rec.ID != 0 { // 可以设置id从1开始
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "当前用户名已注册!",
		})
		return
	}

	pwd := ctx.Query("pwd")     // 密码
	rePwd := ctx.Query("rePwd") // 确认密码
	if pwd != rePwd {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"err":  "密码不一致，请重新确认！",
		})
		return
	}
	user := &models.UserBasic{}
	user.Name = name
	/*
		密码要加密保存
	*/
	salt := fmt.Sprintf("%06d", rand.Int31()) // %06d 整数输出 宽度是6位 不足左边补数字0
	encodePwd := utils.MakeSaltPwd(pwd, salt)
	user.Pwd = encodePwd
	user.Salt = salt
	// 创建用户
	models.CreateUser(user)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
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
		"code":    0,
		"message": "用户删除成功",
	})
}

// UpdateUser
// @Summary 更新用户资料
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "用户名"
// @param pwd  formData string false "密码"
// @param phone formData string false "电话号码"
// @param email  formData string false "邮箱"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(ctx *gin.Context) {
	// 解析id
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	name := ctx.PostForm("name")
	pwd := ctx.PostForm("pwd")
	phone := ctx.PostForm("phone")
	email := ctx.PostForm("email")

	var user models.UserBasic
	user.ID = uint(id)
	user.Name = name
	user.Pwd = pwd
	user.Phone = phone
	user.Email = email

	/*
		邮箱和电话号码的更新需要做校验
		需要符合规范
		这里用到了 govalidator 这个包
	*/
	_, err := valid.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  -1,
			"error": "参数非法,请确认后重试！",
		})
		return
	}

	// 更新用户数据
	models.UpdateUser(user)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "用户数据更新成功",
	})
}
