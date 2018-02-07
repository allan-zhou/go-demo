package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("len(s): %d  cap(s): %d \n", len(s), cap(s))

	/*
	* 以下代码会报错： panic: runtime error: index out of range
	* 
	* s[4] = 4
	* fmt.Printf("len(s): %d \ncap(s): %d \n", len(s), cap(s))
	*/

	// 因为s 是一个slice，所以修改s1，s也同时变化
	s1 := s
	s1[0] = 999
	fmt.Printf("s[0]: %d  , s1[0]: %d \n", s[0], s1[0])
	fmt.Printf("&s: %v  , &s1: %v \n", &s, &s1)


	// make 创建的是一个数组slice切片，适用于 引用类型
	fmt.Println()
	s2 := make([]int, 2, 5)
	fmt.Printf("len(s2): %d  cap(s2): %d \n", len(s2), cap(s2))
	s2[0] = 1
	s2[1] = 2
	s22 := s2
	s22[0] = 99
	fmt.Printf("s2: %v  ,  s22: %v \n", s2, s22)
	fmt.Printf("&s2: %v  ,  &s22: %v \n", &s2, &s22)
	fmt.Printf("*s2: %v  ,  *s22: %v \n", *&s2, *&s22)

	// new 创建的是一个数组常量，适用于 值类型
	fmt.Println()
	s3 := new([3]int)
	s3[0] = 1
	s3[1] = 2
	s33 := *s3
	s33[0] = 99
	fmt.Printf("s3: %v  , s33: %v \n", s3, s33)
	fmt.Printf("&s3: %v  , &s33: %v \n", &s3, &s33)
	fmt.Printf("*&s3: %v  , *&s33: %v \n", *&s3, *&s33)
	fmt.Printf("*s3: %v  \n", *s3)

	fmt.Println()
	s4 := make([]int, 0)
	// s4[0] = 1  报错：panic: runtime error: index out of range
	fmt.Printf("len(s4): %d  cap(s4): %d \n", len(s4), cap(s4))


	// ... 声明的数组不是一个切片
	fmt.Println()
	s5 := [...]int{1, 2, 3}
	fmt.Printf("len(s5): %d  cap(s5): %d \n", len(s5), cap(s5))	
	s55 := s5
	s55[0] = 999
	fmt.Printf("s5: %d  , s55: %d \n", s5, s55)
	fmt.Printf("&s5: %v  , &s55: %v \n", &s5, &s55)
	fmt.Printf("*&s5: %d  , *&s55: %d \n", *&s5, *&s55)
}
