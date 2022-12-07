package utils

import (
	"github.com/spf13/viper"
)

var Instance *Config

type Config struct {
	// 数据库配置
	DB struct {
		Url          string `yaml:"Url"`
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
