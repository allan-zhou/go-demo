package main

import (
	"fmt"
)

func main()  {
	fmt.Println("start")
	
	panic("a server error occurred!")
	
	fmt.Println("end")
}