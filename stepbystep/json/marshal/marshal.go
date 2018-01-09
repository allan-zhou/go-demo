package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Person struct {
	Username string `json:"username"`
	Age      int
	Gender   bool
}

type KV struct {
	Key string `json:"key"`
	Value []byte `json:"value"`
}

func main() {
	p1 := Person{"zhangsan", 20, true}
	var p2 Person

	bytes, _ := json.Marshal(p1)
	// fmt.Print(user1.username)
	// fmt.Print(v)
	fmt.Printf("json format p1: %c \n", bytes)

	json.Unmarshal(bytes, &p2)
	bytes2, _ := json.Marshal(bytes)
	fmt.Printf("json format p2: %c \n", bytes2)

	persons :=[]Person{
		Person{"zhangsan", 20, true},
		Person{"lisi", 21, true},
		Person{"wangwu", 22, false},
	}

	kvs := []KV{}
	for i,p := range persons {
		pAsBytes,_ := json.Marshal(p)

		key:= strconv.Itoa(i)
		kv := KV{ key, pAsBytes}
		kvs = append(kvs, kv)
	}

	str := "["
	fmt.Printf("kvs length is: %v \n",len(kvs))
	for _,kv := range kvs {		
		// fmt.Printf("kv: %v \n", kv)
		str += "{\"key\":"
		str += "\""
		str += kv.Key
		str += "\","
		str += "\"value\":"
		str += string(kv.Value)
		str += "},"
	}
	str = str[:len(str)-1]
	str += "]"
	fmt.Printf("%s \n", str)
}
