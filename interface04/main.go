package main

import "fmt"

type Address struct {
	Name  string
	Phone int
}

// golang中空接口和类型断言使用细节

func main() {
	var userinfo = make(map[string]interface{})
	userinfo["username"] = "张三"
	userinfo["age"] = 20
	userinfo["hobby"] = []string{"睡觉", "吃饭"}

	//fmt.Println(userinfo["hobby"][1])			(map index expression of type interface{})

	var address = Address{
		Name:  "李四",
		Phone: 1632631841,
	}
	userinfo["address"] = address
	fmt.Println(userinfo["address"])

	//var name = userinfo["address"].Name		 userinfo["address"].Name undefined (type interface{} has no field or method Name)
	//fmt.Println(name)

	hobby2, _ := userinfo["hobby"].([]string)
	fmt.Println(hobby2[0])

	address2, _ := userinfo["address"].(Address)
	fmt.Println(address2.Name)

}
