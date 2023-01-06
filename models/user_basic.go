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
	Name          string `form:"name"`
	Pwd           string `form:"pwd"`
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"` // govalidator 直接提供了一个email的校验接口
	Identity      string // 唯一身份标识
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time // 登陆时间
	HeartbeatTime time.Time // 心跳
	LoginOutTime  time.Time `gorm:"column:login_out_time"` // 退出时间 可以通过gorm标签来指定生成的字段名
	DeviceInfo    string    // 设备信息
	Salt          string    // 密码加密盐值
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

// FindUserByNameAndPwd 根据用户名和密码找到用户 后期可以提供 电话号码&邮箱之类的校验方式
func FindUserByNameAndPwd(name, pwd string) (UserBasic, string) {
	// 状态控制
	msg := ""
	// 根据用户名找到salt
	user := FindUserByName(name)
	if user.ID == 0 {
		msg = "用户不存在！"
		return user, msg
	}
	// 将登录密码加密之后，与库中密码做对比
	isPwdRight := utils.ValidSaltPwd(pwd, user.Salt, user.Pwd)
	if !isPwdRight {
		/*
			后续还可以扩充校验次数 间隔时间之类的
		*/
		msg = "密码与当前用户不符！" // 安全性期间不要告诉他用户名和密码哪个不对
		return UserBasic{}, msg
	} else {
		msg = "密码核验通过"
		return user, msg
	}
}

// FindUserByName 通过用户名定位到某个用户
func FindUserByName(name string) UserBasic {
	var user UserBasic
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

// FindUserByTel 通过手机号定位到缪个用户
func FindUserByTel(tel string) UserBasic {
	var user UserBasic
	utils.DB.Where("phone = ?", tel).First(user)
	return user
}

// FindUserByEmail 通过手机号定位到缪个用户
func FindUserByEmail(email string) UserBasic {
	var user UserBasic
	utils.DB.Where("email = ?", email).First(user)
	return user
}

// CreateUser 创建用户
func CreateUser(user *UserBasic) {
	utils.DB.Create(user)
}

// DeleteUser 删除用户
func DeleteUser(user *UserBasic) {
	// 软删除 给用户的相关表字段增加了一个删除时间
	utils.DB.Delete(&user)
}

// UpdateUser 修改用户资料
func UpdateUser(user UserBasic) {
	utils.DB.Model(&user).Updates(UserBasic{
		Name:  user.Name,
		Pwd:   user.Pwd,
		Phone: user.Phone,
		Email: user.Email,
	})
}
