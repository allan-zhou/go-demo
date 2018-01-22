package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func  writefile()  {
	file,_ := os.OpenFile("data.txt", os.O_CREATE | os.O_WRONLY, 0666)
	defer file.Close()	

	writer := bufio.NewWriter(file)
	for index := 0; index < 10; index++ {
		writer.WriteString("hi go~~\r\n")
	}
	writer.Flush()
}

func readfile()  {
	file, _ := os.Open("data.txt")

	reader := bufio.NewReader(file)	
	for {
		result, err := reader.ReadString('\n')

		fmt.Printf(result)

		if err == io.EOF {
			return
		}
	}
}

func main() {

	writefile()
	readfile()
}
