package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Info4 struct {
	gorm.Model
	Money  int
	Dog4ID uint
}
type Dog4 struct {
	gorm.Model
	Name      string
	Infos4    Info4
	GirlGod4s []GirlGod4 `gorm:"many2many:dog4_girl4_gods;"`
}

type GirlGod4 struct {
	gorm.Model
	Name  string
	Dogs4 []Dog4 `gorm:"many2many:dog4_girl4_gods;"`
}

func many2Many() {
	GLOBAL_DB.AutoMigrate(&Dog4{}, &GirlGod4{}, &Info4{})
	//i := Info4{
	//	//Money: 20000,
	//	Money: 200,
	//}
	g1 := GirlGod4{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "琪琪",
	}
	g2 := GirlGod4{
		Model: gorm.Model{
			ID: 2,
		},
		Name: "米米",
	}
	//d1 := Dog4{
	//	Name:      "汪汪1号",
	//	GirlGod4s: []GirlGod4{g1, g2},
	//	Infos4:    i,
	//}
	//d2 := Dog4{
	//	Name:      "汪汪2号",
	//	GirlGod4s: []GirlGod4{g1, g2},
	//	Infos4:    i,
	//}

	//GLOBAL_DB.Create(&d2)

	d := Dog4{
		Model: gorm.Model{
			ID: 1,
		},
	}

	GLOBAL_DB.Preload("GirlGod4s").Preload("Infos4").Find(&d)
	fmt.Println(d)

	var girls []GirlGod4
	GLOBAL_DB.Model(&d).Association("GirlGod4s").Find(&girls)
	GLOBAL_DB.Model(&d).Preload("Dogs4", func(db *gorm.DB) *gorm.DB {
		return db.Joins("Infos4").Where("money<?", 10000)
	}).Association("GirlGod4s").Find(&girls)

	d1 := Dog4{
		Model: gorm.Model{
			ID: 3,
		},
	}
	GLOBAL_DB.Model(&d1).Association("GirlGod4s").Append(&g1, &g2)
	GLOBAL_DB.Model(&d1).Association("GirlGod4s").Delete(&g1)
	GLOBAL_DB.Model(&d1).Association("GirlGod4s").Replace(&g1)
	GLOBAL_DB.Model(&d1).Association("GirlGod4s").Clear()

	fmt.Println(girls)
}
