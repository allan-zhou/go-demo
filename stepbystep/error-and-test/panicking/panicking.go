package main

import (
	"fmt"
)

func func1()  {
	fmt.Println("do something 1")
	defer fmt.Println("end do something 1")
	
	func2()
}

func func2()  {
	fmt.Println("do something 2")
	defer fmt.Println("end do something 2")

	// panic("error occurred in func2")

	func3()
}

func func3()  {
	fmt.Println("do something 3")
	defer fmt.Println("end do something 3")

	panic("error occurred in func3")
}

func main()  {
	func1()
}