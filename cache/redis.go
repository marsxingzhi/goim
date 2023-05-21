package cache

import (
	"fmt"

	"github.com/marsxingzhi/goim/config"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func init() {
	// 不能在这里调用，因为如果在这里调用的话，可能config还未初始化呢
	// NewRedisClient()
}

func NewRedisClient() {
	options := redis.Options{
		Addr:     config.Conf.GetRedisAddr(),     // config.Conf.Redis.Addr
		Password: config.Conf.GetRedisPassword(), // no password set
		DB:       config.Conf.GetDBName(),        // db
	}

	c := redis.NewClient(&options)
	RedisClient = c

	fmt.Println("[redis] connect to redis successfully...")
}
