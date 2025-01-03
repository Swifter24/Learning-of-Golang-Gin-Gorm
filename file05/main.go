package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("file05/test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString("hello world\r\n") //写入缓存
	write.Flush()                        //缓存内容写入文件
}
