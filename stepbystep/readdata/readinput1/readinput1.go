package main

import (
	"fmt"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "12.34 / 56 / go"
	format                 = "%f / %d / %s"
)

func main() {
	fmt.Println("please enter your name:")
	fmt.Scanln(&firstName, &lastName)
	//fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("your full name is: %s %s \n", firstName, lastName)

	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("read the result is: ", f, i, s)
}
