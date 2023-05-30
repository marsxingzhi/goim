package config

import (
	"fmt"
	"io/ioutil"

	conf "github.com/marsxingzhi/goim/pkg/config"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Name       string      `yam:"name"`
	ServerID   int         `yaml:"server_id"`
	Etcd       *conf.Etcd  `yaml:"etcd"`
	Redis      *conf.Redis `yaml:"redis"`
	Mysql      *conf.Mysql `yaml:"mysql"`
	GrpcServer *conf.Grpc  `yaml:"grpc_server"` // auth是grpc服务，因此肯定得有grpc配置
}

var config = new(Config)

func GetConfig() *Config {
	return config
}

func init() {
	bytes, err := ioutil.ReadFile("/Users/geyan/go/src/goim/configs/auth.yaml")

	if err != nil {
		fmt.Println("[config] failed to read yaml file: ", err)
		panic(err)
	}

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		fmt.Println("[config] failed to unmarshal bytes: ", err)
		panic(err)
	}

	fmt.Printf("[config] init auth cofnig successfully\nconfig: %+v\n", config)
	fmt.Printf("Etcd: %+v\n", config.Etcd)
	fmt.Printf("Redis: %+v\n", config.Redis)
	fmt.Printf("Mysql: %+v\n", config.Mysql)
	fmt.Printf("GrpcServer: %+v\n", config.GrpcServer)
}
