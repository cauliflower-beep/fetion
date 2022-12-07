package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name          string
	Pwd           string
	Phone         string
	Email         string
	Identity      string // 唯一身份标识
	ClientIp      string
	ClientPort    string
	LoginTime     uint64 // 登陆时间
	HeartbeatTime uint64 // 心跳
	LogOutTime    uint64 // 退出时间
	DeviceInfo    string // 设备信息
}

func (table *UserBasic) TableName() string {
	return "t_user_basic"
}
