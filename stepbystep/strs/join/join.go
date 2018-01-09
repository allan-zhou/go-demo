package main

import (
	"fmt"
	"strings"
	"bytes"
	"strconv"
)

func main() {
	// 1.
	str := "hello"
	str += "go"
	fmt.Printf("str1: \n%s \n\n", str)

	//2.
	str2 := "hello go hi china"
	str2AsArr := strings.Split(str2, " ")
	str2 = strings.Join(str2AsArr, " ")
	fmt.Printf("str2: \n%s \n\n", str2)

	//3. 
	var buffer bytes.Buffer
	for i := 0; i < 100; i++ {
		buffer.WriteString("index:")
		buffer.WriteString(strconv.Itoa(i))
		buffer.WriteString(",")
	}
	fmt.Printf("str3: \n%v \n\n", buffer.String())

}
