package xzmysql

import (
	"fmt"
	"time"

	"github.com/marsxingzhi/goim/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

/**
mysql的封装
*/

var cli *MysqlClient

type MysqlClient struct {
	db *gorm.DB
}

func NewMysqlClient(conf *config.Mysql) *MysqlClient {
	cli = &MysqlClient{}
	db, err := connectDb(conf)
	if err != nil {
		fmt.Println("[xzmysql] failed to new mysql client: ", err)
		panic(err)
	}
	cli.db = db
	return cli
}

// pathWrite := strings.Join([]string{MySql.DbUser, ":", MySql.DbPassword, "@tcp(", MySql.DbHost, ":", MySql.DbPort, ")/", MySql.DbName, "?charset=utf8mb4&parseTime=true"}, "")
func connectDb(conf *config.Mysql) (*gorm.DB, error) {
	var (
		pathWrite string
		sources   = make([]gorm.Dialector, 0)
		replicas  = make([]gorm.Dialector, 0)
	)
	// dsn - %s:%s@(%s)/%s?charset=utf8&parseTime=true&loc=Local
	for i, v := range conf.Sources {
		path := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True", v.UserName, v.Password, v.Addr, v.Db)
		if i == 0 {
			pathWrite = path
			continue
		}
		sources = append(sources, mysql.Open(path))
	}

	for _, v := range conf.Replicas {
		path := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True", v.UserName, v.Password, v.Addr, v.Db)
		replicas = append(replicas, mysql.Open(path))
	}

	db, err := gorm.Open(mysql.Open(pathWrite), &gorm.Config{
		SkipDefaultTransaction: false, // 禁用默认事务(true: Error 1295: This command is not supported in the prepared statement protocol yet)
		PrepareStmt:            false, // 创建并缓存预编译语句(true: Error 1295)
	})
	if err != nil {
		fmt.Println("[xzmysql] failed to gorm open: ", err)
		return nil, err
	}

	db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  sources,
		Replicas: replicas,
		// sources/replicas 负载均衡策略
		Policy: dbresolver.RandomPolicy{},
	}).SetMaxIdleConns(conf.MaxIdleConns).
		SetMaxOpenConns(conf.MaxOpenConns).
		SetConnMaxLifetime(time.Duration(conf.MaxLifetime) * time.Millisecond))

	return db, nil
}
