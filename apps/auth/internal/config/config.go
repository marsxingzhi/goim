package config

import (
	"fmt"
	"io/ioutil"

	conf "github.com/marsxingzhi/goim/pkg/config"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Name     string      `yam:"name"`
	ServerID int         `yaml:"server_id"`
	Etcd     *conf.Etcd  `yaml:"etcd"`
	Redis    *conf.Redis `yaml:"redis"`
	Mysql    *conf.Mysql `yaml:"mysql"`
}

var config = new(Config)

func GetConfig() *Config {
	return config
}

func init() {
	bytes, err := ioutil.ReadFile("./configs/auth.yaml")

	if err != nil {
		fmt.Println("[config] failed to read yaml file: ", err)
		panic(err)
	}

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		fmt.Println("[config] failed to unmarshal bytes: ", err)
		panic(err)
	}

	fmt.Printf("[config] init auth cofnig successfully, and config: %+v\n", config)
}
