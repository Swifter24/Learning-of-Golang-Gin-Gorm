package main

func TestUpdate() {
	//update	只更新你选择的字段

	//dbRes := GLOBAL_DB.Model(&TestUser{}).Where("name=?", "zwz").Update("name", "zwz946")
	//fmt.Println(dbRes.RowsAffected)

	//updates	只更新所有字段，一种为map 一种为结构体 结构体0值不参与更新

	//map
	var user1 []TestUser
	GLOBAL_DB.Where("name = ?", "第一条").Find(&user1).Updates(map[string]interface{}{
		"Name": "第二条",
		"Age":  0,
	})
	//结构体
	//var user2 []TestUser
	//GLOBAL_DB.Where("name = ?", "rh").Find(&user2).Updates(TestUser{Name: "第一条", Age: 0})

	//save	无论如何都更新包括0值

	// var user []TestUser
	// dbRes := GLOBAL_DB.Model(&TestUser{}).Where("name=?", "zwz946").Find(&user)
	// for k := range user {
	// 	user[k].Age = 23
	// }
	// dbRes.Save(&user)

}
