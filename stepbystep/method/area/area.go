package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{width: 4, height: 5}
	r2 := Rectangle{4, 5}
	c1 := Circle{radius: 4}
	c2 := Circle{4}

	fmt.Println("r1.area", r1.area())
	fmt.Println("r2.area", r2.area())
	fmt.Println("c1.area", c1.area())
	fmt.Println("c2.area", c2.area())
}
