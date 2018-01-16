package main

import (
	"fmt"
	"sync"
	my "zhoujl/go-demo/stepbystep/packages/mypkg"
)

type Info struct{
	mu sync.Mutex
	
	// todo
	phone string
}

func update(info *Info)  {
	info.mu.Lock()

	info.phone = "123"

	info.mu.Unlock()
}

func main()  {
	fmt.Println(my.GetName())
}