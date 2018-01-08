package main

import (
	"fmt"
)

func main() {
	arr := []int{11, 22, 33, 44, 55}

	fmt.Println("For")
	for i := 0; i < len(arr); i++ {
		fmt.Println(i)
	}

	fmt.Println("For range")
	for a, b := range arr {
		fmt.Printf(" index: %d " , a)
		fmt.Printf(" value: %d " , b)
		fmt.Println()
	}
}
