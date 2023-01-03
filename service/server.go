package service

import "github.com/gin-gonic/gin"

// GetIndex
// @Tags 首页
// @Success 200 {string} pong
// @Router /index [get]
func GetIndex(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
