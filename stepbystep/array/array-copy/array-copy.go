package main

import (
	"fmt"
)

func main()  {
	arr1 := new([5]int)
	arr2 := [5]int{11,22,33,44,55}
	fmt.Printf("arr1[0]: %d  arr1[1]: %d \n", arr1[0], arr1[1])
	fmt.Printf("arr2[0]: %d  arr2[1]: %d \n", arr2[0], arr2[1])

	// 把arr1赋值给arr2
	// 需要在做一次数组内存的拷贝操作	
	arr1[0] = 1
	arr2 = *arr1
	fmt.Println()
	fmt.Printf("arr1[0]: %d  arr1[1]: %d \n", arr1[0], arr1[1])
	fmt.Printf("arr2[0]: %d  arr2[1]: %d \n", arr2[0], arr2[1])

	//修改arr2
	arr2[0] = 999
	fmt.Println()
	fmt.Printf("arr1[0]: %d  arr1[1]: %d \n", arr1[0], arr1[1])
	fmt.Printf("arr2[0]: %d  arr2[1]: %d \n", arr2[0], arr2[1])
	
}
