package test

import (
	"fetion/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

// TestConn 测试Gorm链接
func TestConn(t *testing.T) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:admin123@tcp(192.168.7.114:3306)/fetion?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect mysql")
	}

	// 迁移 shema
	_ = db.AutoMigrate(&models.UserBasic{}) // 如果没有该表 会依据实体类去创建

	// create
	user := &models.UserBasic{}
	user.Name = "lufy"
	db.Create(user)

	// Read
	db.First(&user, 1)                  // 根据整型主键查找
	db.First(&user, "Name = ?", "lufy") // 查找 Name 字段值为 lufy 的记录

	// Update - 将 user 的 phone 更新为 15706290582
	db.Model(&user).Update("Phone", 15706290582)
	// Update - 更新多个字段
	db.Model(&user).Updates(models.UserBasic{Name: "conan", Pwd: "F42"}) // 仅更新非零值字段
	db.Model(&user).Updates(map[string]interface{}{"Name": "Conan", "Pwd": "F43"})

	// Delete - 删除 user
	db.Delete(&user, 1)

}
