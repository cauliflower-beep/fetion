package main

import (
	"fetion/models"
	"fetion/router"
	"fetion/utils"
	"fmt"
)

func main() {
	// 配置初始化
	_serverConf := utils.InitConfig("./config")
	fmt.Println(_serverConf)

	// 初始化数据库连接
	utils.InitMysql(_serverConf.DB.Dns)
	// 初始化表
	// 表迁移 保持 scheme 为最新 应对添加表字段的情况
	_ = utils.DB.AutoMigrate(&models.UserBasic{})  // 自动迁移用户表
	_ = utils.DB.AutoMigrate(&models.Message{})    // 自动迁移消息表
	_ = utils.DB.AutoMigrate(&models.Relation{})   // 自动迁移用户表
	_ = utils.DB.AutoMigrate(&models.GroupBasic{}) // 自动迁移消息表

	// 初始化redis连接
	utils.InitRedis(_serverConf)

	r := router.Router()
	_ = r.Run(":9080") // 默认监听在本机 8080 端口 http://127.0.0.1:8080/index
}
