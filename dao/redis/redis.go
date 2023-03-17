package redis

import (
	"fmt"
	"web_app/settings"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(config *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			config.Host,
			config.Port,
		),
		Password: config.Password, // 密码
		DB:       config.Db,       // 数据库
		PoolSize: config.PoolSize, // 连接池大小
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		fmt.Printf("redis connect failed, err:%v\n", err)
		zap.L().Error("redis connect failed", zap.Error(err))
	}
	fmt.Printf("redis connect success\n")
	zap.L().Info("redis connect success")
	return
}

func Close() {
	_ = rdb.Close()
}
