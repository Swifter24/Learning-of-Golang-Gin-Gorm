package main

import (
	"fmt"
	"gorm.io/gorm"
)

//hasOne

type Dog1 struct {
	gorm.Model
	Name       string
	GirlGod1ID uint
}

type GirlGod1 struct {
	gorm.Model
	Name string
	Dog1 Dog1
}

func One2oneHasOne() {
	//d := Dog1{
	//	Name: "小冉",
	//}
	//g := GirlGod1{
	//	Name: "宋豪",
	//	Dog1: d,
	//}

	//GLOBAL_DB.Create(&g)

	var girl1 GirlGod1
	GLOBAL_DB.Preload("Dog1").Find(&girl1, 2) //预加载查询
	fmt.Println(girl1)
	GLOBAL_DB.AutoMigrate(&GirlGod1{}, &Dog1{})


}
