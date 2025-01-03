package main

import (
	"fmt"
	"reflect"
)

// v.Elem().SetInt()	反射修改值
func reflectFn(x interface{}) {
	v := reflect.ValueOf(x)
	//b, _ := x.(*int64)		//类型断言也可以修改
	//*b = 123
	fmt.Println(v.Kind(), v.Elem().Kind())
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(123)
	}
}

func main() {
	var a int64 = 120

	reflectFn(&a)
	fmt.Println(a)

}
