package main

import (
	"fmt"
	"gorm.io/gorm"
)

//belongsTo

type Dog struct {
	gorm.Model
	Name      string
	GirlGodID uint
	GirlGod   GirlGod
}

type GirlGod struct {
	gorm.Model
	Name string
}

func One2oneBelongsTo() {
	g := GirlGod{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "小红",
	}
	g2 := GirlGod{
		Model: gorm.Model{
			ID: 2,
		},
		Name: "小绿",
	}
	d := Dog{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "小明",
	}

	//GLOBAL_DB.AutoMigrate(&Dog{}) //belongs to 关系下 外键关联的主表在创建外键表时会同时创建主表
	//GLOBAL_DB.Create(&d)
	//GLOBAL_DB.Create(&g)

	//关联操作同时适用于hasone
	GLOBAL_DB.Model(&d).Association("GirlGod").Append(&g)   // 使用 Append 来更新关联
	GLOBAL_DB.Model(&d).Association("GirlGod").Replace(&g2) // 使用 Replace 来更改关联
	GLOBAL_DB.Model(&d).Association("GirlGod").Delete(&g2)  // 使用 Delete 来删除关联
	GLOBAL_DB.Model(&d).Association("GirlGod").Clear()      // 使用 Clear 来清楚关联

	var dog Dog
	GLOBAL_DB.Preload("GirlGod").Find(&dog, 1) //预加载查询
	fmt.Println(dog)
}
