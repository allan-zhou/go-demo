package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string  `json:"name"`
	Addr   string  `json:"addr,omitempty"`
	Age    int     `json:"age,string"`
	Salary float64 `json:"-"`
	Status bool    `json:"status,string"`
}

func main() {
	u1 := User{
		Name: "zhangsan", Addr: "", Age: 20, Salary: 12345, Status: true,
	}

	u1AsBytes, _ := json.Marshal(u1)

	fmt.Printf("json format : %s \n", string(u1AsBytes))

}
