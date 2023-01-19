package models

import "gorm.io/gorm"

// GroupBasic 群信息
type GroupBasic struct {
	gorm.Model
	Name     string // 群名
	LeaderId int64  // 群主
	Icon     string // 群头像
	Level    string // 群等级
	Desc     string // 描述
}

func (table *GroupBasic) TableName() string {
	return "t_group_basic"
}
