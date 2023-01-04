package main

import (
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

	r := router.Router()
	_ = r.Run(":9080") // 默认监听在本机 8080 端口 http://127.0.0.1:8080/index
}
