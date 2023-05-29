package xzredis

import (
	"context"
	"fmt"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/marsxingzhi/goim/pkg/config"
	"github.com/redis/go-redis/v9"
)

/**
redis的封装
*/

var (
	Cli *RedisClient
)

type RedisClient struct {
	Client   *redis.ClusterClient
	RedsSync *redsync.Redsync
}

func NewRedisClient(conf *config.Redis) *redis.ClusterClient {
	// 1. 创建集群redis
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: conf.Address,
	})

	// 2. 判断是否能够链接到redis
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("[xzredis] failed to ping redis: ", err)
		panic(err)
	}

	// 3. redis锁
	rsPool := goredis.NewPool(client)
	redisSync := redsync.New(rsPool)

	Cli = &RedisClient{
		Client:   client,
		RedsSync: redisSync,
	}
	return client
}
