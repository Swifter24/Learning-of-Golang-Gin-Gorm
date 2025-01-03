package main

import "fmt"

func main() {
	var ch1 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}
	close(ch1)             //for range记得关闭管道，普通for可以不关闭
	for val := range ch1 { //注意管道没有key
		fmt.Println(val)
	}
}
