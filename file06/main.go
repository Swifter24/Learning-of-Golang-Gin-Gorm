package main

import "io/ioutil"

func main() {
	str := "hello golang"
	err := ioutil.WriteFile("file06/test.txt", []byte(str), 0666)
	if err != nil {
		panic(err)
		return
	}
}
