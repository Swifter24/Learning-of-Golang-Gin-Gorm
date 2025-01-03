package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Info struct {
	gorm.Model
	Money  int
	Dog3ID uint
}
type Dog3 struct {
	gorm.Model
	Name       string
	GirlGod3ID uint
	Infos      Info
}

type GirlGod3 struct {
	gorm.Model
	Name  string
	Dogs3 []Dog3
}

func hasMany() {
	GLOBAL_DB.AutoMigrate(&Dog3{}, &GirlGod3{}, &Info{})
	//d1 := Dog3{
	//	Model: gorm.Model{
	//		ID: 1,
	//	},
	//	Name: "汪汪1号",
	//}
	//d2 := Dog3{
	//	Model: gorm.Model{
	//		ID: 2,
	//	},
	//	Name: "汪汪2号",
	//}
	//g := GirlGod3{
	//	Model: gorm.Model{
	//		ID: 1,
	//	},
	//	Name:  "琪琪",
	//	Dogs3: []Dog3{d1, d2},
	//}
	//
	//GLOBAL_DB.Create(&g)

	var girl GirlGod3

	GLOBAL_DB.Preload("Dogs3").First(&girl)
	GLOBAL_DB.Preload("Dogs3", "name=?", "汪汪1号").First(&girl) //带条件的预加载
	GLOBAL_DB.Preload("Dogs3", func(db *gorm.DB) *gorm.DB {
		return db.Where("name=?", "汪汪2号")
	}).First(&girl) //内嵌函数预加载

	GLOBAL_DB.Preload("Dogs3.Infos", "money>100").Preload("Dogs3", "name=?", "汪汪1号").First(&girl) //带条件的预加载
	GLOBAL_DB.Preload("Dogs3.Infos", "money>100").Preload("Dogs3").First(&girl)                   //带条件的预加载 所有的条件都会带出来
	GLOBAL_DB.Preload("Dogs3", func(db *gorm.DB) *gorm.DB {
		return db.Joins("Infos").Where("money>?", 500)
	}).First(&girl) //带条件的预加载 Joins只带出符合条件的

	fmt.Println("Girl hasMany:", girl)
}
