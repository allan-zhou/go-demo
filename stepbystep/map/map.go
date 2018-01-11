package main

import (
	"fmt"
)

func main() {
	maplist := map[int]string{1: "one", 2: "two", 3: "three"}

	fmt.Printf("maplist. length is: %d \n", len(maplist))
	for key, value := range maplist {
		fmt.Printf("key:%d  value:%s \n", key, value)
	}

	maplistCopy := maplist
	maplistCopy[1] = "i am one"
	fmt.Printf("maplist modified. length is: %d \n", len(maplist))
	for key, value := range maplist {
		fmt.Printf("key:%d  value:%s \n", key, value)
	}

	mapFunc := map[string]func() int{
		"f1": func() int { return 1 },
		"f2": func() int { return 2 },
	}

	fmt.Printf("mapFunc modified. length is: %d \n", len(mapFunc))
	for key, value := range mapFunc {
		fmt.Printf("key:%s  value:%v \n", key, value)
	}

	fmt.Println()
	// 判断 map[key] 是否存在
	if mapValue,ok := maplist[999]; ok {
		fmt.Printf("maplist[999] is exist, and value is: %s \n", mapValue)
	} else {
		fmt.Printf("maplist[999] is not exist \n")
	}
	
	if mapValue,ok := maplist[3]; ok {
		fmt.Printf("maplist[3] is exist, and value is: %s \n", mapValue)

		delete(maplist, 3)
		fmt.Printf("deleted maplist[3] \n")
		fmt.Printf("now, maplist[3]==\"\" is: %v", maplist[3]=="")
	} else {
		fmt.Printf("maplist[3] is not exist \n")
	}

	
}
