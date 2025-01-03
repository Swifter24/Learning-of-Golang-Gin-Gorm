package main

import "fmt"

// 接口是一个规范

type Usber interface {
	start()
	stop()
}

// 如果接口里面有方法的话，必须要通过结构体或者通过自定义类型实现这个接口

type Computer struct {
}

type Phone struct {
	Name string
}
type Camera struct {
}

// 手机要实现usb接口的话必须要实现usb接口中的所有方法
func (p *Phone) start() {
	fmt.Println(p.Name, "启动")
}
func (p *Phone) stop() {
	fmt.Println(p.Name, "关机")
}
func (c *Camera) start() {
	fmt.Println("相机启动")
}
func (c *Camera) stop() {
	fmt.Println("相机关机")
}
func (c *Computer) work(usb Usber) {
	if _, ok := usb.(*Phone); ok { //类型断言
		usb.start()
	} else {
		usb.stop()
	}
}

func main() {
	p := &Phone{
		Name: "华为手机",
	}
	c := &Camera{}
	com := &Computer{}

	//var p1 Usber //golang中接口就是一个数据类型
	//var c1 Usber
	//p1 = p //表示手机实现usb接口
	//c1 = c
	//p1.start()
	//p1.stop()
	//c1.start()
	//c1.stop()

	com.work(p)
	com.work(c)
}
