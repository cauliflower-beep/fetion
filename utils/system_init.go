package utils

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// Instance 定义成全局变量 方便后续调用
var Instance *Config

var DB *gorm.DB

type Config struct {
	// 数据库配置
	DB struct {
		Dns          string `yaml:"Dns"`
		MaxIdleConns int    `yaml:"MaxIdleConns"`
		MaxOpenConns int    `yaml:"MaxOpenConns"`
	} `yaml:"DB"`
}

// InitConfig 配置初始化
func InitConfig(filePath string) *Config {
	conf := viper.New()
	conf.AddConfigPath(filePath)
	conf.SetConfigName("app")
	conf.SetConfigType("yaml")
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}

	// 将配置文件内容映射到 struct 中
	if err := conf.Unmarshal(&Instance); err != nil {
		panic(err)
	}

	return Instance
}

func InitMysql(dns string) {
	// 自定义日志模板 打印sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 终端打印
		logger.Config{
			SlowThreshold: time.Second, // 慢查询阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // caise
		},
	)
	DB, _ = gorm.Open(mysql.Open(dns), &gorm.Config{Logger: newLogger})
}
