package main

import (
	"io"
	"os"
)

func copy(srcFileName string, dstFileName string) (err error) {
	sFile, err1 := os.Open(srcFileName)
	defer sFile.Close()
	dFile, err2 := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0666)
	defer dFile.Close()
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}

	var tempSlice = make([]byte, 128)
	for {
		n1, e1 := sFile.Read(tempSlice)
		if e1 == io.EOF {
			break
		}
		if e1 != nil {
			return e1
		}
		_, e2 := dFile.Write(tempSlice[:n1])
		if e2 != nil {
			return e2
		}
	}
	return nil
}

func main() {
	copy("file07_copy/test.txt", "file08_copy/test.txt")
}
