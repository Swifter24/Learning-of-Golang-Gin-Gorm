package main

type Jiazi struct {
	ID          uint
	Name        string
	Xiaofengche []Xiaofengche `gorm:"polymorphic:Owner;polymorphicValue:huhu"`
}
type Yujie struct {
	ID          uint
	Name        string
	Xiaofengche Xiaofengche `gorm:"polymorphic:Owner;polymorphicValue:abaaba"`
}
type Xiaofengche struct {
	ID        uint
	Name      string
	OwnerType string
	OwnerID   uint
}

func Polymorphic() {
	GLOBAL_DB.AutoMigrate(&Jiazi{}, &Yujie{}, &Xiaofengche{})
	GLOBAL_DB.Create(&Jiazi{Name: "夹子", Xiaofengche: []Xiaofengche{ //多态不能被别人同时拥有，但同时可以拥有多个多态
		{Name: "小风车1"},
		{Name: "小风车2"},
	}})
	GLOBAL_DB.Create(&Yujie{Name: "御姐", Xiaofengche: Xiaofengche{
		Name: "大风车",
	}})
}
