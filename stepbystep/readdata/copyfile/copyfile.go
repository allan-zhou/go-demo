package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func CopyFile(srcFileName, destFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer srcFile.Close()

	destFile, err := os.OpenFile(destFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer destFile.Close()

	return io.Copy(destFile, srcFile)
}

func main() {
	CopyFile("src.txt", "dest.txt")
	fmt.Println("copy completed")
}
