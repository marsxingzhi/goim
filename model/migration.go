package model

import "fmt"

// 数据迁移
func migration() {
	// https://gorm.io/zh_CN/docs/migration.html
	err := GormDB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{})
	if err != nil {
		fmt.Println("[model] failed to auto migrate: ", err)
		panic(err)
	}
}
