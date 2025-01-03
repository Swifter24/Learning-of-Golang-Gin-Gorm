package main

import (
	"fmt"
	"reflect"
)

// TypeOf()类型  Name()类型名称	Kind()底层的具体类型
func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("%v %v %v\n", v, v.Name(), v.Kind())
}

type Myint int
type Person struct {
	Name string
}

func main() {
	a := 1
	b := "1"
	c := true
	d := 1
	var e Myint = 1
	var p = Person{
		Name: "zhangsan",
	}
	f := make([]int, 3, 6)
	g := make(map[int]int, 3)
	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(&d)
	reflectFn(e)
	reflectFn(f)
	reflectFn(g)
	reflectFn(p)

}
