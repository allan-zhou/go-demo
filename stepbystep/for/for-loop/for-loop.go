package main

import (
	"fmt"
)

func main() {
	count := 10

	// 1
	for i := 0; i < count; i++ {
		fmt.Printf("%v ", i)
	}

	fmt.Println()
	
	// 2
	i := 0
START:
	fmt.Printf("%v ", i)
	i++
	if i < count {
		goto START
	}
}
