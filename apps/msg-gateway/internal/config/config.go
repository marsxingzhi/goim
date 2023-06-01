package config

import (
	"fmt"
	conf "github.com/marsxingzhi/goim/pkg/config"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Name     string `yaml:"name"`
	ServerID int    `yaml:"server_id"`
	// grpc服务配置
	GrpcServer *conf.Grpc       `yaml:"grpc_server"`
	MsgServer  *conf.GrpcServer `yaml:"msg_server"`
	Etcd       *conf.Etcd       `yaml:"etcd"`
	Redis      *conf.Redis      `yaml:"redis"`
	WsServer   *conf.WsServer   `yaml:"ws_server"`
}

var config = new(Config)

func GetConfig() *Config {
	return config
}

func init() {
	bytes, err := ioutil.ReadFile("/Users/geyan/go/src/goim/configs/msg_gateway.yaml")

	if err != nil {
		fmt.Println("[config] failed to read yaml file: ", err)
		panic(err)
	}

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		fmt.Println("[config] failed to unmarshal bytes: ", err)
		panic(err)
	}

	fmt.Println("[config] init msg_gateway config successfully")
}
