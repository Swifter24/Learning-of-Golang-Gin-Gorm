package main

import "fmt"

func TestDelete() {
	var user TestUser
	//GLOBAL_DB.First(&user).Delete(&user)
	GLOBAL_DB.Unscoped().First(&user).Delete(&user) //硬删除

	//SQL原生操作
	var raw []TestUser
	GLOBAL_DB.Raw("Select * from test_users where name = ?", "zwz946").Scan(&raw)
	fmt.Println(raw)
}
