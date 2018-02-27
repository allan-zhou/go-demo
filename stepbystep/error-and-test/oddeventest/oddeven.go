package main

import (
	"fmt"
	"go-demo/stepbystep/error-and-test/oddeventest/even"
)

func main()  {
	fmt.Println("odd and even testing")

	for i := 0; i < 100; i++ {
		fmt.Printf("%d is a Even? %t \n",i, even.Even(i))
	}
}