package test

import (
	"fetion/models"
	"fetion/utils"
	"testing"
)

// TestGetUserList 获取用户列表 go test -run TestGetUserList -v
func TestGetUserList(t *testing.T) {
	// 初始化数据库连接
	dsn := "root:admin123@tcp(127.0.0.1:3306)/fetion?charset=utf8&parseTime=True&loc=Local"
	utils.InitMysql(dsn)

	models.GetUsersList()
}
