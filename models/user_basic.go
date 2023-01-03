package models

import (
	"fetion/utils"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// UserBasic 用户信息
type UserBasic struct {
	gorm.Model
	Name          string
	Pwd           string
	Phone         string
	Email         string
	Identity      string // 唯一身份标识
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time // 登陆时间
	HeartbeatTime time.Time // 心跳
	LoginOutTime  time.Time `gorm:"column:login_out_time"` // 退出时间 可以通过gorm标签来指定生成的字段名
	DeviceInfo    string    // 设备信息
}

func (table *UserBasic) TableName() string {
	return "t_user_basic"
}

// GetUserList 获取全部用户
func GetUserList() []*UserBasic {
	userList := make([]*UserBasic, 10)
	utils.DB.Find(&userList)
	for _, user := range userList {
		fmt.Println(user.Name)
	}
	return userList
}
