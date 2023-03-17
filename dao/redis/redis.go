package redis

import (
	"fmt"

	"github.com/spf13/viper"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port"),
		),
		Password: viper.GetString("redis.password"), // 密码
		DB:       viper.GetInt("redis.db"),          // 数据库
		PoolSize: viper.GetInt("redis.pool_size"),   // 连接池大小
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
