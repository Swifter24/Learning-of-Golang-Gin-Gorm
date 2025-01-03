package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1() 你好golang", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数器减一
}
func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2() 你好golang", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数器减一
}
func main() {
	wg.Add(1)  //协程计数器加一
	go test1() //开启一个协程	主死从随
	wg.Add(1)  //协程计数器加一
	go test2() //开启一个协程	主死从随
	for i := 0; i < 10; i++ {
		fmt.Println("main() 你好golang", i)
		time.Sleep(time.Millisecond * 50)
	}
	//time.Sleep(time.Second)
	wg.Wait()
	fmt.Println("主线程退出")
}
