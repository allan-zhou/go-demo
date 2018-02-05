package main

import (
	"fmt"
	"math/big"
)

func main()  {
	i := new(big.Int)
	i.SetString("10", 16) // octal
	fmt.Println(i)
	fmt.Printf("%X \n",i)

	i2 := new(big.Int)
	i2.SetString("23952764745190639315453386127001698040085062535617457724687902786301481890590", 16) // hex
	fmt.Println(i2)
}