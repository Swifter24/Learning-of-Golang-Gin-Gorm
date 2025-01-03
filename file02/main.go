package main

import (
	"bufio"
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
	reader := bufio.NewReader(file)
	var fileStr string
	for {
		str, err := reader.ReadString('\n') //表示一次读取一行
		if err == io.EOF {
			fileStr += str //记得加上
			break
		}
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fileStr += str
	}

	fmt.Println(fileStr)

}
