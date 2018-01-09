package main

import (
	"fmt"
)

type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
}

func main() {
	str := "Go is a nice language"
	fmt.Printf("string length is: %v \n", len(str))
	for i, v := range str {
		fmt.Printf("index: %v   value: %c \n", i, v)
	}

	cars := []Car{
		Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "zhangsan"},
		Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "lisi"},
		Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "wangwu"},
		Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "zhangsan"},
		Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "lisi"},
		Car{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "wangwu"},
		Car{Make: "Chery", Model: "S22L", Colour: "white", Owner: "zhangsan"},
		Car{Make: "Fiat", Model: "Punto", Colour: "violet", Owner: "lisi"},
	}

	for i, car := range cars {
		fmt.Printf("index: %v   value: %v \n", i, car)
	}
}
