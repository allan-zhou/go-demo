package main

import (
	"fmt"
	"go-demo/stepbystep/error-and-test/panic_package/parse"
	// "./parse"
)

func main()  {
	var examples = []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}

	for _,v := range examples {
		fmt.Printf("pasring %s \n", v)

		numbsers,err := parse.Parse(v)
		if err != nil {
			fmt.Println(err)
			// continue
		}

		fmt.Printf("parsing result: %v \n\n", numbsers)
	}
}