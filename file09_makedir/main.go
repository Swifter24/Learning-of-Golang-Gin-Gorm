package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("file09_makedir/abc.txt", 0666)
	if err != nil {
		fmt.Println(err)
	}
	err = os.MkdirAll("file09_makedir/dir1/dir2/abc.txt", 0666)
	if err != nil {
		fmt.Println(err)
	}
}
