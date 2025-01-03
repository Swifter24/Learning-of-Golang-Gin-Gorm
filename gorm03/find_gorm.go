package main

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type TestUser1 struct {
	Name string
}

func TestFind() {
	//map
	result := map[string]interface{}{}
	GLOBAL_DB.Model(&TestUser{}).First(&result)
	fmt.Println(result)

	//结构体
	var user1 TestUser
	GLOBAL_DB.Model(&TestUser{}).Last(&user1)
	fmt.Println(user1)

	//主键带参数检索
	var user2 TestUser
	dbRes := GLOBAL_DB.Model(&TestUser{}).Last(&user2, 14)      //主键检索
	fmt.Println(errors.Is(dbRes.Error, gorm.ErrRecordNotFound)) //没找到返回true

	//字符串条件检索
	var user3 TestUser
	GLOBAL_DB.Where("name=? AND age=?", "zwz", 18).Model(&TestUser{}).Find(&user3)
	fmt.Println(user3)

	//结构体条件查询
	var user4 TestUser
	GLOBAL_DB.Where(TestUser{Name: "rh", Age: 9}).Model(&TestUser{}).Find(&user4)
	fmt.Println(user4)

	//map条件查询
	var user5 TestUser
	GLOBAL_DB.Where(map[string]interface{}{
		"Name": "sh",
	}).Model(&TestUser{}).Find(&user5)
	fmt.Println(user5)

	//条件叠加or检索
	var user6 TestUser
	GLOBAL_DB.Where("name=? AND age=?", "aaa", 18).Or("name=?", "zzl").Model(&TestUser{}).Find(&user6)
	fmt.Println(user6)

	//条件叠加排序检索
	var user7 []TestUser //切片可以查多条
	GLOBAL_DB.Order("name asc,age").Model(&TestUser{}).Find(&user7)
	fmt.Println(user7)

	//内联条件查询 map
	var user8 TestUser
	GLOBAL_DB.Model(&TestUser{}).Find(&user8, map[string]interface{}{"age": 44})
	fmt.Println(user8)

	//内联条件查询 结构体
	var user9 TestUser
	GLOBAL_DB.Model(&TestUser{}).Find(&user9, TestUser{Age: 44})
	fmt.Println(user9)

	//select选择查询字段  omit也能用
	var user10 TestUser
	GLOBAL_DB.Select("name").Model(&TestUser{}).Find(&user10, TestUser{Age: 44})
	fmt.Println(user10)

	//高级查询
	var u []TestUser1
	GLOBAL_DB.Model(&TestUser{}).Limit(5).Find(&u)
	fmt.Println(u)
}
