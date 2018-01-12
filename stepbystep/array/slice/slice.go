package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("len(s): %d \ncap(s): %d \n", len(s), cap(s))

	/*
	* 以下代码会报错： panic: runtime error: index out of range
	* 
	* s[4] = 4
	* fmt.Printf("len(s): %d \ncap(s): %d \n", len(s), cap(s))
	*/

	s2 := make([]int, 2, 5)
	fmt.Printf("len(s2): %d \ncap(s2): %d \n", len(s2), cap(s2))


	s3 := new([3]int)
	s4 := make([]int, 0)
	fmt.Printf("len(s3): %d \ncap(s3): %d \n", len(*s3), cap(*s3))
	fmt.Printf("len(s4): %d \ncap(s4): %d \n", len(s4), cap(s4))
}
