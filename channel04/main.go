package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func write(ch chan<- int) { //只写
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func read(ch <-chan int) { //只读
	for num := range ch {
		fmt.Printf("%d ", num)
	}
	wg.Done()
}
func main() {
	//ch1 := make(chan<- int, 3) //只写管道
	//ch2 := make(<-chan int, 3) //只读管道

	ch := make(chan int, 10)
	wg.Add(1)
	go write(ch)
	wg.Add(1)
	go read(ch)

	wg.Wait()
	fmt.Println("done")
	fmt.Println("--------------------------------------------------------------------------------")
	s1 := make(chan int, 10)
	s2 := make(chan string, 10)

	for i := 0; i < 10; i++ {
		s1 <- i
	}
	for i := 0; i < 10; i++ {
		s2 <- strconv.Itoa(i)
	}
	for {
		select { //多路复用 并发	使用select不需要关闭channel
		case num := <-s1:
			fmt.Println("int", num)
		case num := <-s2:
			fmt.Println("string", num)
		default:
			fmt.Println("no one")
			return
		}
	}
}
