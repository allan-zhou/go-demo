package main

import (
	"fmt"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	Side float32
}

type Rectangle struct {
	Width float32
	Height float32
}

func (s *Square) Area() float32 {
	return s.Side * s.Side
}

func (r *Rectangle) Area() float32 {
	return r.Width * r.Height
}

func main() {
	s := &Square{5}
	r := new(Rectangle)
	r.Height = 3
	r.Width = 5

	// shape := s
	shapes := []Shaper{s, r}

	for i,_ := range shapes {		
		fmt.Printf("shape details: %v \n", shapes[i])
		fmt.Printf("the shape area is: %v \n", shapes[i].Area())
	}
}
