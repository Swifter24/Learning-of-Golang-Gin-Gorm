package main

import (
	"database/sql"
	"time"
)

type Model struct {
	UUID uint      `gorm:"primaryKey"`
	Time time.Time `gorm:"column:my_time"`
}

type TestUser struct {
	Model        Model   `gorm:"embedded"`
	Name         string  `gorm:"default:zz"`
	Email        *string `gorm:"comment:年龄"`
	Age          uint8   `gorm:"not null"`
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}

func TestUserCreate() {
	GLOBAL_DB.AutoMigrate(&TestUser{})
}
