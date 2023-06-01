package xzredis

import "context"

func MSet(values ...interface{}) error {
	// MSET 是一个原子性(atomic)操作， 所有给定键都会在同一时间内被设置， 不会出现某些键被设置了但是另一些键没有被设置的情况。
	return cli.Client.MSet(context.Background(), values...).Err()
}

func Get(key string) (val string, err error) {
	key = realKey(key)
	val, err = cli.Client.Get(context.Background(), key).Result()
	return
}

func realKey(key string) string {
	if cli != nil {
		return cli.Prefix + key
	}
	return key
}
