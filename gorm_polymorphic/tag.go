package main

type Jiazi2 struct {
	ID           uint
	Name         string
	Xiaofengche2 []Xiaofengche2 `gorm:"foreignkey:Jiazi2Name;references:Name"` //指定外键加重引用

	//Xiaofengche2 []Xiaofengche2 `gorm:"many2many:jiazi_fengche;foreignkey:Name;joinForeignKey:jiazi;references:Name;joinReferences:fengche"`
	//多对多重写，分别是两个表中不同的Name	join是指分别的别名
}

type Xiaofengche2 struct {
	ID   uint
	Name string
	//JiaziID   uint
	Jiazi2Name string
}

func Tag() {
	GLOBAL_DB.AutoMigrate(&Jiazi2{}, &Xiaofengche2{})
	GLOBAL_DB.Create(&Jiazi2{Name: "大夹子", Xiaofengche2: []Xiaofengche2{
		{Name: "小风车1"},
		{Name: "小风车2"},
	}})
}
