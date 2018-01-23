package main

import (
	"fmt"
	"os"
	"flag"
)

// var Newline = flag.Bool("n", false, "print newline")
var loop = flag.Int("a", 1, "loop times")
var name = flag.String("b", "", "your name")

func main()  {
	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags

	var s string = ""
	
	for index := 0; index < flag.NArg(); index++ {
		fmt.Println(flag.Arg(index))
	}

	os.Stdout.WriteString(s)
}