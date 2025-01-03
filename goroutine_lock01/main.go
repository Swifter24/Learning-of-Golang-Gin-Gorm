package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var m = make(map[int]int, 0)

func test(num int) {
	mutex.Lock()
	var sum = 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	m[num] = sum
	fmt.Println(num, m[num])
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}
func main() {
	for i := 0; i < 40; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
}
