package main

import "fmt"

// 一个结构体可以实现多个接口	//也可以进行嵌套
type Animals interface {
	SetName(string)
	GetName() string
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}
func (d Dog) GetName() string {
	return d.Name
}

type Cat struct {
	Name string
}

func (c *Cat) SetName(name string) {
	c.Name = name
}
func (c Cat) GetName() string {
	return c.Name
}

func main() {
	dog := &Dog{
		Name: "小黑",
	}
	cat := &Cat{
		Name: "小花",
	}
	var d1 Animals
	d1 = dog
	var c1 Animals = cat
	fmt.Println(d1.GetName())
	d1.SetName("hua hua")
	fmt.Println(d1.GetName())

	fmt.Println(c1.GetName())
	c1.SetName("la la")
	fmt.Println(c1.GetName())

}
