package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p Person) GetInfo() string {
	var str = fmt.Sprintf("姓名:%v  年龄:%v", p.Name, p.Age)
	return str
}

func (p *Person) SetInfo(name string, age int) {
	p.Name = name
	p.Age = age
}

func (p Person) Print() {
	fmt.Println("这是一个打印方法。。。")
}
func PrintStructField(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}

	field0 := t.Field(0) //第一个字段
	fmt.Println(field0.Name)
	fmt.Println(field0.Type)
	fmt.Println(field0.Tag.Get("json"))

	if field1, ok := t.FieldByName("Age"); ok {
		fmt.Println(field1.Name)
		fmt.Println(field1.Type)
		fmt.Println(field1.Tag.Get("json"))
	}

	fieldCount := t.NumField() //获取有几个字段
	fmt.Println(fieldCount)

	fmt.Println(v.FieldByName("Name")) //获取字段值
	fmt.Println(v.FieldByName("Age"))

}

func PrintStructFn(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}

	method0 := t.Method(0)
	fmt.Println(method0.Name)
	fmt.Println(method0.Type)

	method1, ok := t.MethodByName("Print")
	if ok {
		fmt.Println(method1.Name)
		fmt.Println(method1.Type)
	}

	//调用方法
	v.Method(1).Call(nil)
	v.MethodByName(("Print")).Call(nil)
	info1 := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info1)

	var params []reflect.Value
	params = append(params, reflect.ValueOf("李四"))
	params = append(params, reflect.ValueOf(20))
	v.MethodByName("SetInfo").Call(params)
	info2 := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info2)

	fmt.Println(t.NumMethod())
}
func main() {
	var p = Person{
		Name: "张三",
		Age:  19,
	}
	PrintStructField(p)
	PrintStructFn(&p)
}
