package main

import (
	"fmt"
)

type Phone struct{}

func (p *Phone) Call() string {
	return "a call"
}

type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "a picture"
}

type CameraPhone struct {
	Phone
	Camera
}

func main() {
	cp := new(CameraPhone)

	fmt.Printf("give a call. %s \n", cp.Call())
	fmt.Printf("take a picture. %s \n", cp.TakeAPicture())
}
