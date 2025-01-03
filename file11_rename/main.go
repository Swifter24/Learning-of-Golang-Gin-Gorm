package main

import "os"

func main() {
	err := os.Rename("file11_rename/test.txt", "file11_rename/text.txt")
	if err != nil {
		return
	}
}
