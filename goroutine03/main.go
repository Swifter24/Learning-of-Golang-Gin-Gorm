package main

import (
	"fmt"
	"sync"
	"time"
)

//func main() {
//	start := time.Now().Unix()
//	for i := 2; i < 120000; i++ {
//		var flag = true
//		for j := 2; j < i; j++ {
//			if i%j == 0 {
//				flag = false
//				break
//			}
//		}
//		if flag {
//			fmt.Println(i, "是素数")
//		}
//	}
//	end := time.Now().Unix()
//	fmt.Println(end-start, "秒")
//}

var wg sync.WaitGroup

func test(n int) {
	for i := (n-1)*30000 + 1; i < n*30000; i++ {
		var flag = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i, "是素数")
		}
	}
	wg.Done()
}

func main() {
	start := time.Now().Unix()
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	end := time.Now().Unix()
	fmt.Println("执行完毕", end-start, "秒")
}
