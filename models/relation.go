package models

import "gorm.io/gorm"

// Relation 人员关系
type Relation struct {
	gorm.Model
	Uid      uint // 用户id
	FriendId uint // 好友id
	Type     int  // 关系类型 好友|拉黑
	Desc     string
}

func (table *Relation) TableName() string {
	return "t_rela"
}
