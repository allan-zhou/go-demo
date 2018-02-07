package main

import (
	"fmt"
)

func DoAppendNumbers(base,v []int) []int {
	return append(base, v...)
}

func test1(){
	s1 := []string{"hi","go"}

	s1 = append(s1, "chaoyang", "beijing")
	fmt.Println(s1)

	s2 := []string{"a","b", "c", "d", "e"}
	// 如果s2的长度小于s1，则会截断字符串
	copy(s2, s1[:2])
	fmt.Println(s2)

	// copy s1&s2 to s3
	s3 := make([]string, len(s1) + len(s2))
	copy(s3[:len(s1)], s1)
	copy(s3[len(s1):], s2)
	fmt.Println(s3)
}

func test2(){
	int1 := []int{1,2,3}
	int2 := []int{4,5,6}

	ret := DoAppendNumbers(int2, int1)
	fmt.Println(ret)
}

func main(){
	test1()

	test2()
}