package models

import (
	"gorm.io/gorm"
)

// Message 用户信息
type Message struct {
	gorm.Model
	From     uint   // 发送用户id
	To       uint   // 接收用户id
	ChatType string // 聊天类型 1-私聊 2-群聊 3-广播
	MsgType  int    // 消息类型 1-文本 2-图片 3-语音 4-视频
	Content  string // 消息内容
	Pic      string // 图片内容
	Url      string // 链接内容
	Desc     string // 消息描述
	Amount   int    // 其他数字内容 发送频率|文件大小等
}

func (table *Message) TableName() string {
	return "t_msg"
}
