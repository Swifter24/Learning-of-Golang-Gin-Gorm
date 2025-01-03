package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//1、打开文件
	file, err := os.Open("file01/main.go")
	defer file.Close()
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	//2、读取文件
	var strSlice []byte
	var tempSlice = make([]byte, 128)

	for {
		n, err := file.Read(tempSlice)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		strSlice = append(strSlice, tempSlice[:n]...)
	}

	fmt.Println(string(strSlice))

}
