package service

import (
	"fetion/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUsers
// @Tags 获取用户列表
// @Success 200 {string} json{"code","message"}
// @Router /getUsers [get]
func GetUsers(ctx *gin.Context) {
	users := models.GetUserList()
	ctx.JSON(http.StatusOK, gin.H{
		"name": users,
	})
}
