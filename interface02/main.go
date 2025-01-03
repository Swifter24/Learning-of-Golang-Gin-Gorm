package main

import "fmt"

type A interface{} //空接口 表示没有任何约束  任意的类型都可以实现空接口

func show(a interface{}) { //空接口可以直接当做类型来使用，可以表示任何类型
	fmt.Printf("%T\n", a)
}

func Print(x interface{}) {
	switch x.(type) {
	case string:
		fmt.Println("string")
		break
	case int:
		fmt.Println("int")
		break
	default:
		fmt.Println("error")
	}
}
func main() {
	var b interface{} //空接口可以直接当做类型来使用，可以表示任何类型
	var a A
	var str = "hello"
	var m1 = make(map[string]interface{}) //空接口可以直接当做类型来使用，可以表示任何类型
	var m2 = make([]interface{}, 2, 4)    //空接口可以直接当做类型来使用，可以表示任何类型
	a = str                               //让字符串实现A这个接口
	fmt.Printf("%v %T\n", a, a)

	b = 20
	fmt.Printf("%v %T\n", b, b)

	show(true)

	m1["username"] = "zhangsan"
	m1["age"] = 18
	m2[0] = "lisi"
	m2[1] = 19
	fmt.Println(m1)
	fmt.Println(m2)

	//类型断言
	var s interface{}
	s = "hello golang"
	v, ok := s.(string)
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println(v)
	}

	Print("hello")
}
