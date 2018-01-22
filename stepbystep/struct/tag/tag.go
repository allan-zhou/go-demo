package main

import (
	"fmt"
	"reflect"
)

type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func main() {
	tt := &TagType{true, "jialiang", 100}

	for index := 0; index < 3; index++ {
		reflectTag(tt, index)
	}
}

func reflectTag(tt *TagType, ix int) {
	ttType := reflect.TypeOf(*tt)
	ixFiled := ttType.Field(ix)
	fmt.Printf("%v \n", ixFiled.Tag)

	fmt.Printf("%v \n", *tt)
	fmt.Printf("%v \n", tt)
}
