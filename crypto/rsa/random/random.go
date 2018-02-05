package main

import (
	"fmt"
	// "crypto/rand"
	// "io"
	// "math/big"
	"math/rand"
)

func main()  {
	r := rand.New(rand.NewSource(100))	
	fmt.Println(r.Int())
	fmt.Println(r.Intn(100))
	fmt.Println(r.Perm(5))
}