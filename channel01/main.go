package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 10
	ch <- 11
	<-ch
	x := <-ch
	fmt.Println(x)                                           //11
	fmt.Printf("%v  %T  %v  %v\n", ch, ch, cap(ch), len(ch)) //0xc0000a4080  chan int  3  0

	//管道阻塞：放不下数据或者取不出值了
}
