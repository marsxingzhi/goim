package model

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/marsxingzhi/goim/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var GormDB *gorm.DB

func NewMySQL() {
	c := config.Conf
	dsn := strings.Join([]string{c.GetMySQLUser(), ":", c.GetMySQLPassword(), "@tcp(", c.GetMySQLHost(), ":", strconv.Itoa(c.GetMySQLPort()), ")/", c.GetMySQLDBName(), "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	fmt.Println("[mysql] dsn: ", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{ // 命名策略
			SingularTable: true, // 表单数化，不要加s
		},
	})
	if err != nil {
		fmt.Println("[mysql] failed to init mysql: ", err)
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("[mysql] failed to getSqlDB: ", err)
		panic(err)
	}
	// 用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	GormDB = db

	migration()

	fmt.Println("[mysql] connect to mysql successfully...")
}

var MongoClient *mongo.Client

// mongodb
func NewMongoDB() {
	// 设置客户端连接配置
	c := config.Conf
	uri := strings.Join([]string{"mongodb://", c.GetMongoDBAddr(), ":", strconv.Itoa(c.GetMongoDBPort())}, "")
	fmt.Println("[mongoDB] uri: ", uri)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println("[mongoDB] failed to connect mongoDB: ", err)
		panic(err)
	}

	// 检查链接
	if err := client.Ping(context.TODO(), nil); err != nil {
		fmt.Println("[mongoDB] failed to ping mongoDB: ", err)
		panic(err)
	}

	MongoClient = client
	fmt.Println("[mongoDB] connect to mongoDB successfully...")
}
