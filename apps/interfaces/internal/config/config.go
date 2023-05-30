package config

import (
	"fmt"
	"io/ioutil"

	conf "github.com/marsxingzhi/goim/pkg/config"
	"gopkg.in/yaml.v3"
)

// 配置
type Config struct {
	Name     string      `yam:"name"`
	ServerID int         `yaml:"server_id"`
	Port     int         `yaml:"port"`
	Etcd     *conf.Etcd  `yaml:"etcd"`
	Redis    *conf.Redis `yaml:"redis"`
}

var config = new(Config)

func GetConfig() *Config {
	return config
}

func init() {
	bytes, err := ioutil.ReadFile("./configs/api_gateway.yaml")

	if err != nil {
		fmt.Println("[config] failed to read yaml file: ", err)
		panic(err)
	}

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		fmt.Println("[config] failed to unmarshal bytes: ", err)
		panic(err)
	}

	fmt.Printf("[config] init successfully, and config: %+v\n", config)
}
