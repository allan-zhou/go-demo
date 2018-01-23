package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "go "
	if len(os.Args) > 1 {
		fmt.Println(os.Args)
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("hello ", who)
}
