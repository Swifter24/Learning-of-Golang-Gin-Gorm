package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("file04/test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString("WriteString\r\n")

	var file1 = "Write"
	file.Write([]byte(file1))
}
