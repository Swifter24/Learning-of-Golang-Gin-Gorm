package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

type CInfo struct {
	Name string
	Age  int
}

func (c CInfo) Value() (driver.Value, error) {
	str, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return string(str), nil
}
func (c *CInfo) Scan(value interface{}) error {
	str, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(str, c)
}

type Array []string

func (a Array) Value() (driver.Value, error) {
	if len(a) > 0 {
		var str string = a[0]
		for _, v := range a[1:] {
			str = str + "," + v
		}
		return str, nil
	} else {
		return "", nil
	}
}
func (a *Array) Scan(value interface{}) error {
	str, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	*a = strings.Split(string(str), ",")
	return nil
}

type CUser struct {
	ID    uint
	Info  CInfo //`gorm:"text"` 自定义
	Array Array
}

func Customize() {
	GLOBAL_DB.AutoMigrate(&CUser{})
	GLOBAL_DB.Create(&CUser{Info: CInfo{Name: "zz", Age: 23}, Array: Array{"1", "2", "3"}})

	var u CUser
	GLOBAL_DB.First(&u)
	fmt.Println(u)
}
