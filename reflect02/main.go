package main

import (
	"fmt"
	"reflect"
)

// ValueOf()	Kind()底层的具体类型
func reflectFn(x interface{}) {
	v := reflect.ValueOf(x)
	//v,_:=x.(int)	类型断言
	//var m = v+10
	kind := v.Kind()	//获取类型
	switch kind {	//类型匹配
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println(v.Int() + 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Println(v.Uint() + 10)
	case reflect.Float32, reflect.Float64:
		fmt.Println(v.Float() + 10)
	case reflect.String:
		fmt.Println(v.String() + "10")
	default:
		fmt.Println("Unknow type!")
	}

}

func main() {
	a := 1
	b := "1"
	c := -1
	d := 1.1

	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)
}
