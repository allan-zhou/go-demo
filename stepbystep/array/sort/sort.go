package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{12, 4, 8, 54}

	// 排序
	sort.Ints(ints)
	fmt.Printf("sort.Ints(ints) \n")
	if sort.IntsAreSorted(ints) {
		fmt.Printf("ints is sorted \n")
		for _, value := range ints {
			fmt.Printf("value: %d  ", value)
		}

		// 搜索
		x := 1
		index := sort.SearchInts(ints, x)
		fmt.Println(index)
		x = 5
		index = sort.SearchInts(ints, x)
		fmt.Println(index)
		x = 33
		index = sort.SearchInts(ints, x)
		fmt.Println(index)

		// if index > 0 {
		// 	fmt.Printf("\n\n %d is in []ints, and index is: %d", x, index)
		// } else {
		// 	fmt.Printf("\n %d is not in []ints, and index is: %d ", x, index)
		// }
	}
}
