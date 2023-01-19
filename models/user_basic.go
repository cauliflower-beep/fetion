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
	Uid           int64
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
	if user.ID == 0 { // 要求用户表id不从0开始
		msg = "用户不存在！"
		return user, msg
	}
	// 将登录密码加密之后，与库中密码做对比
	isPwdRight := utils.ValidSaltPwd(pwd, user.Salt, user.Pwd)
	if !isPwdRight {
		/*
			后续还可以扩充校验次数 间隔时间之类的
		*/
		msg = "密码与当前用户不符！" // 安全性起见不要告诉他用户名和密码哪个不对
		return UserBasic{}, msg
	} else {
		msg = "密码核验通过"
		/*
			token引入：
			客户端频繁向服务器请求数据，服务端频繁的去数据库查询用户名和密码进行对比，判断用户名和密码正确与否，并作出相应提示。
			在这样的背景下，token便应运而生。
			token定义：
			本质是服务端生成的一串字符串，以作客户端进行请求的一个令牌。
			当第一次登录后，服务器生成一个token并返回给客户端，以后客户端只需带上这个token前来请求数据即可，无需再次带上用户名和密码。
			使用token的目的：
			减轻服务器压力，减少频繁的数据库查询，使服务器更加健壮。

			如何使用token？
			1.用设备号/设备mac地址作为token（推荐）
			  客户端在登录的时候获取设备的设备号/mac地址，并将其作为参数传递到服务端。
			  服务端收到参数之后，将其作为token保存在数据库，并将该token设置到session中，客户端每次请求的时候同意拦截，
			  并将客户端传递的token和服务器端session中的token做对比，相同则放行，不同则拒绝。
			  优缺点：
			  客户端和服务器端统一了一个唯一的标识Token，而且保证了每一个设备拥有了一个唯一的会话。
			  缺点是客户端需要带设备号/mac地址作为参数传递，而且服务器端还需要保存；
			  优点是客户端不需重新登录，只要登录一次以后一直可以使用，至于超时的问题是由服务器这边来处理。
			  如何处理？若服务器的Token超时后，服务器只需将客户端传递的Token向数据库中查询，
			  同时并赋值给变量Token，如此，Token的超时又重新计时。
			2.用session值作为token
			  客户端：只需携带用户名和密码登陆即可。
			  服务器：接收到用户名和密码后并判断，如果正确了就将本地获取sessionID作为Token返回给客户端，
			  客户端以后只需带上请求数据即可。
			  这种方式使用的好处是方便，不用存储数据；
			  缺点就是当session过期后，客户端必须重新登录才能进行访问数据。
		*/
		// token加密 将当前时间进行md5加密，作为token返回给客户端
		token := utils.Md5Encode(fmt.Sprintf("%d", time.Now().Unix()))
		utils.DB.Model(&UserBasic{}).Where("id = ?", user.ID).Update("identity", token)
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
