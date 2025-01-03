package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("file10_remove/abc.txt")
	if err != nil {
		fmt.Println(err)
	}
	err = os.RemoveAll("file10_remove/dir1/")
	if err != nil {
		fmt.Println(err)
	}
}
