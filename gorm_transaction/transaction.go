package main

type TMG struct {
	ID   uint
	Name string
}

func TestTransaction() {
	GLOBAL_DB.AutoMigrate(&TMG{})

	//flag := false
	//GLOBAL_DB.Transaction(func(tx *gorm.DB) error {
	//	tx.Create(&TMG{Name: "Golang"})
	//	tx.Create(&TMG{Name: "Java"})
	//	tx.Create(&TMG{Name: "Python"})
	//	tx.Transaction(func(tx *gorm.DB) error {	//内层事务返回error不执行
	//		tx.Create(&TMG{Name: "Go"})
	//		return errors.New("error")
	//	})
	//	fmt.Println("已执行")
	//	if flag {
	//		return nil
	//	}else {
	//		return errors.New("error")
	//	}
	//})

	//手动回滚
	tx := GLOBAL_DB.Begin() //开始事务
	tx.Create(&TMG{Name: "JS"})
	tx.Create(&TMG{Name: "TS"})
	if true {
		tx.Rollback() //回滚后无效果，保持原样
	}
	tx.Commit()                   //提交事务
	tx.Create(&TMG{Name: "HTML"}) //事务已经结束，此条不会执行

	//tx.SavePoint("tiao")
	//tx.RollbackTo("tiao")
}
