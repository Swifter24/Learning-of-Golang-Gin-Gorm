package main

import (
	"fmt"
	"io/ioutil"
)

func copy(strFileName string, dstFileName string) (err error) {
	byteStr, err := ioutil.ReadFile(strFileName)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dstFileName, byteStr, 0666)
	if err != nil {
		return err
	}
	fmt.Println("复制文件成功")
	return nil
}

func main() {
	copy("file06/test.txt", "file07_copy/test.txt")
}
