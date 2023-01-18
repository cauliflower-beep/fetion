package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
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

var (
	ctx = context.Background()
	DB  *gorm.DB
	RDB *redis.Client
)

type Config struct {
	// 数据库配置
	DB struct {
		Dns          string `yaml:"Dns"`
		MaxIdleConns int    `yaml:"MaxIdleConns"`
		MaxOpenConns int    `yaml:"MaxOpenConns"`
	} `yaml:"DB"`

	// redis配置
	redis struct {
		Addr        string `yaml:"addr"`
		Pwd         string `yaml:"pwd"`
		DB          int    `yaml:"db"`
		PoolSize    int    `yaml:"poolSize"`
		MinIdleConn int    `yaml:"minIdleConn"`
	}
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
			Colorful:      true,        // 彩色
		},
	)
	DB, _ = gorm.Open(mysql.Open(dns), &gorm.Config{Logger: newLogger})
}

// InitRedis 初始化redis连接
func InitRedis(conf *Config) {
	RDB = redis.NewClient(&redis.Options{
		Addr:         conf.redis.Addr,
		Password:     conf.redis.Pwd,
		DB:           conf.redis.DB,
		PoolSize:     conf.redis.PoolSize,
		MinIdleConns: conf.redis.MinIdleConn,
	})
	// 测试是否链接成功
	pong, err := RDB.Ping(ctx).Result()
	fmt.Println(pong, err)
}

const (
	PublishKey = "websocket" // 管道名
)

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("{Publish.msg}|", msg)
	err = RDB.Publish(ctx, channel, msg).Err()
	return err
}

// Subscribe 订阅Redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := RDB.Subscribe(ctx, channel)
	fmt.Println("{subscribe}|", sub)
	msg, err := sub.ReceiveMessage(ctx)
	return msg.Payload, err
}
