package main

import (
	"fmt"
)

func main() {
	str := "hello"

	fmt.Printf("str string len: %d \n", len(str))
	for i, value := range str {
		fmt.Printf("index: %d  value(char): %c  value(int): %d \n", i, value, value)
	} 

	str2 := []byte(str)
	fmt.Printf("\n str2 []byte len: %d \n", len(str2))
	for i, value := range str2 {
		fmt.Printf("index: %d  value(char): %c  value(int): %d \n", i, value, value)
	}
	
	// copy
	str3 := make([]byte,10)
	copyCount := copy(str3, str)
	fmt.Printf("\n copy str string to str3 []byte, count elements:%d", copyCount)
	fmt.Printf("\n str3 len: %d \n", len(str3))
	for i,value := range str3 {
		fmt.Printf("index: %d  value(char): %c  value(int): %d \n", i, value, value)
	}

	// append
	str3 = append(str3, 'A', 'B')
	fmt.Printf("\n append to str3 []byte with 'A'、'B'")
	fmt.Printf("\n str3 len: %d \n", len(str3))
	for i,value := range str3 {
		fmt.Printf("index: %d  value(char): %c  value(int): %d \n", i, value, value)
	}
	
	str3 = append(str3, str2...)
	fmt.Printf("\n append to str3 []byte with str2[] byte")
	fmt.Printf("\n str3 len: %d \n", len(str3))
	for i,value := range str3 {
		fmt.Printf("index: %d  value(char): %c  value(int): %d \n", i, value, value)
	}

	// 截取字符串
	str4 := str[0:2]
	fmt.Printf("\n substr \n str4 len: %d \n str4 value: %s \n", len(str4), str4)

	// 修改字符串
	c := []byte(str4)
	c[1] = 'i'
	str4 = string(c)

	fmt.Printf("\n modify str \n str4 len: %d \n str4 value: %s \n", len(str4), str4)

}
