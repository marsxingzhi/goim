package main

import (
	"fmt"

	"github.com/marsxingzhi/goim/cache"
	"github.com/marsxingzhi/goim/config"
	"github.com/marsxingzhi/goim/model"
)

func init() {
	config.Init()
	cache.NewRedisClient()
	model.NewMySQL()
	model.NewMongoDB()
}

func main() {
	fmt.Println("Hello World")
}
