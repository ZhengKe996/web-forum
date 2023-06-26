package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"go-web-example/settings"
)

var rdb *redis.Client

func InitDB(config *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password, // 密码
		DB:       config.DB,       // 数据库
		PoolSize: config.PoolSize, // 连接池大小
	})
	if _, err := rdb.Ping().Result(); err != nil {
		return err
	}
	return
}

func Close() {
	rdb.Close()
}
