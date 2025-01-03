package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GLOBAL_DB *gorm.DB

func main() {
	//dst := "root:123456@tcp(127.0.0.1:3306)/zwz?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:123456@tcp(127.0.0.1:3306)/zwz?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 171,
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true, //逻辑外键（在代码里自动体现外键关系）
	})
	if err != nil {
		fmt.Println(err)
	}
	//连接池
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)           //连接池中最大的空闲连接数
	sqlDB.SetMaxOpenConns(100)          //连接池最大的连接数量
	sqlDB.SetConnMaxLifetime(time.Hour) //连接池中连接最大可复用时间

	GLOBAL_DB = db

	//Polymorphic()
	Tag()
}
