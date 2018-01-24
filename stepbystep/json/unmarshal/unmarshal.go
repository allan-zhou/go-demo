package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string
	addr   string 
	age    int
	salary float64
}

func main() {
	// 事先知晓被解析的JSON数据的结构
	var u1 User
	str := `{"Name":"zhangsan","Addr":"beijing","Age":20,"Salary":12345}`
	json.Unmarshal([]byte(str), &u1)
	fmt.Println(u1)

	// 事先不知晓被解析的JSON数据的结构
	var data interface{}
	str2 :=`{"a":"a", "b":"BB", "c":["c","CC"], "d":[{"d1":"d","d2":"dd"},{"d1":"d","d2":"dd"}]}`
	json.Unmarshal([]byte(str2), &data)
	fmt.Println(data)

	if m,ok := data.(map[string]interface{}); ok {
		for k,v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case float64:
				fmt.Println(k,"is float64",vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}
	}
}
