package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	byte, err := ioutil.ReadFile("file01/main.go")
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(byte))
}
