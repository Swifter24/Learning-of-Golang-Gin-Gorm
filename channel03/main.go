package main

import (
	"fmt"
	"sync"
	"time"
)

const Mul = 32

var wg sync.WaitGroup

func putNum(intChan chan int) {
	for i := 2; i < 1200000; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for val := range intChan {
		var flag = true
		for i := 2; i < val; i++ {
			if val%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- val
		}
	}
	//close(primeChan)	//如果一个channel关闭了就没法给这个channel发送数据了
	exitChan <- true
	wg.Done()
}
func printPrime(primeChan chan int) {
	for prime := range primeChan {
		fmt.Println(prime)
	}
	wg.Done()
}
func main() {
	intChan := make(chan int, 10000)
	primeChan := make(chan int, 50000)
	exitChan := make(chan bool, Mul)

	start := time.Now().Unix()
	wg.Add(1)
	go putNum(intChan)

	for i := 1; i <= Mul; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}
	wg.Add(1)
	go printPrime(primeChan)

	wg.Add(1)
	go func() {
		for i := 0; i < Mul; i++ {
			<-exitChan
		}
		//关闭primeChan
		close(primeChan)
		wg.Done()
	}()
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println("执行完毕", end-start)
}
