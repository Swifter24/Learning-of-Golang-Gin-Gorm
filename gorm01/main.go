package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	Name string
}

func main() {
	//dst := "root:123456@tcp(127.0.0.1:3306)/zwz?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:123456@tcp(127.0.0.1:3306)/zwz?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 171,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "gav_",
			//SingularTable: true,
		},
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

	//迁移
	_ = db.AutoMigrate(&User{})
	//M := db.Migrator()
	//M.CreateTable(&User{})
	//fmt.Println(M.HasTable(&User{}))	//是否存在表 或者查表名 gav_Users
	//_ = M.DropTable(&User{})	//删除表
	//_ = M.RenameTable(&User{},"gav_user_two")	//重命名表

	fmt.Println(db)
}
