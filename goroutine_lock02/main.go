package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.RWMutex

func write() {
	mutex.Lock()
	fmt.Println("write")
	time.Sleep(time.Second * 2)
	mutex.Unlock()
	wg.Done()
}
func read() {
	mutex.RLock() //并行执行
	fmt.Println("read")
	time.Sleep(time.Second * 2)
	mutex.RUnlock()
	wg.Done()
}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
}
