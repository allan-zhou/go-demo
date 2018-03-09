package main

import (
	"fmt"
)

func badcall()  {
	panic("bad call")
}

func test()  {
	defer func ()  {
		if e:= recover(); e != nil {
			fmt.Printf("Panicing %s\n", e)
		}
	}()

	badcall()

	fmt.Println("after bad call")
}

func main()  {
	fmt.Println("start")

	test()

	fmt.Println("end")
}