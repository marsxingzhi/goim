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
	cli *RedisClient
)

type RedisClient struct {
	Client   *redis.Client
	RedsSync *redsync.Redsync
	Prefix   string // 当前服务器上的redis加一个前缀
}

func NewRedisClient(conf *config.Redis) *redis.Client {
	// 1. 创建集群redis
	//client := redis.NewClusterClient(&redis.ClusterOptions{
	//	Addrs: conf.Address,
	//})

	// 单机redis
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Address[0],
		Password: conf.Password,
		DB:       conf.Db,
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

	cli = &RedisClient{
		Client:   client,
		RedsSync: redisSync,
		Prefix:   conf.Prefix,
	}
	return client
}
