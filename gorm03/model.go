package main

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	UUID uint      `gorm:"primaryKey"`
	Time time.Time `gorm:"column:my_time"`
}

type TestUser struct {
	Name string `gorm:"default:zz"`
	Age  uint8  `gorm:"not null"`
	gorm.Model
}

func TestUserCreate() {
	GLOBAL_DB.AutoMigrate(&TestUser{})

}
