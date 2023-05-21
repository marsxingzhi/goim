package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// 1. 属性Server一定要大写，否则无法赋值成功
// 2. tag标签后面的server可以加双引号，也可以不加，甚至只加一个，也能读取成功
type Config struct {
	Server  Server  `yaml:"server"`
	Redis   Redis   `yaml:"redis"`
	MySQL   MySQL   `yaml:"mysql"`
	MongoDB MongoDB `yaml:"mongodb"`
}

type Server struct {
	AppMode string `yaml:"appmode"`
	Port    int    `yaml:"port"`
}

type Redis struct {
	DB       string `yaml:"db"`
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DBName   int    `yaml:"db_name"`
}

type MySQL struct {
	DB       string `yaml:"db"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

type MongoDB struct {
	DBName   string `yaml:"db_name"`
	Addr     string `yaml:"addr"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

var Conf *Config

func (c *Config) GetRedisAddr() string {
	return c.Redis.Addr
}

func (c *Config) GetRedisPassword() string {
	return c.Redis.Password
}

func (c *Config) GetDBName() int {
	return c.Redis.DBName
}

func (c *Config) GetMySQLHost() string {
	return c.MySQL.Host
}

func (c *Config) GetMySQLPort() int {
	return c.MySQL.Port
}

func (c *Config) GetMySQLUser() string {
	return c.MySQL.User
}

func (c *Config) GetMySQLPassword() string {
	return c.MySQL.Password
}

func (c *Config) GetMySQLDBName() string {
	return c.MySQL.DBName
}

func (c *Config) GetMongoDBName() string {
	return c.MongoDB.DBName
}

func (c *Config) GetMongoDBAddr() string {
	return c.MongoDB.Addr
}

func (c *Config) GetMongoDBPassword() string {
	return c.MongoDB.Password
}

func (c *Config) GetMongoDBPort() int {
	return c.MongoDB.Port
}

func Init() {
	bytes, err := ioutil.ReadFile("config/config.yaml")

	if err != nil {
		fmt.Println("[config] failed to read yaml file: ", err)
		panic(err)
	}

	if err := yaml.Unmarshal(bytes, &Conf); err != nil {
		fmt.Println("[config] failed to unmarshal bytes: ", err)
		panic(err)
	}

	fmt.Printf("[config] init %+v successfully...\n", Conf)
}
