package main

import "fmt"

func CreatedTest() {
	// dbres := GLOBAL_DB.Create(&TestUser{Name: "zwz", Age: 18})	//常规创建
	// fmt.Println(dbres.Error, dbres.RowsAffected)
	// if dbres.Error != nil {
	// 	fmt.Println("创建失败")
	// } else {
	// 	fmt.Println("创建成功")
	// }

	dbres := GLOBAL_DB.Create(&[]TestUser{ //用切片创建
		{Name: "zwz", Age: 18},
		{Name: "sh", Age: 10},
		{Name: "rh", Age: 9},
		{Name: "wzr", Age: 8},
		{Name: "sh", Age: 72},
		{Name: "zzl", Age: 10},
	})
	fmt.Println(dbres.Error, dbres.RowsAffected)
	if dbres.Error != nil {
		fmt.Println("创建失败")
	} else {
		fmt.Println("创建成功")
	}

	// dbres := GLOBAL_DB.Select("Name").Create(&TestUser{Name: "zwz", Age: 18})	//只想创建Name
	// dbres := GLOBAL_DB.Omit("Name").Create(&TestUser{Name: "zwz", Age: 18})	//只想除去Name
	// fmt.Println(dbres.Error, dbres.RowsAffected)
	// if dbres.Error != nil {
	// 	fmt.Println("创建失败")
	// } else {
	// 	fmt.Println("创建成功")
	// }
}
