package main

import (
	"fmt"
	"errors"
)

var (
	err1 = errors.New("my custom err one")
)

func main()  {
	fmt.Println("error test")
	fmt.Printf("Error: %s \n", err1)
}