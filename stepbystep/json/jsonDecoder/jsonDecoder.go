package main

import (
	"encoding/json"
	"log"
	"os"
	"fmt"
)

type User struct {
	Name string
	Age  int
	Addr string
}

func main() {
	file, _ := os.Open("../jsonEncoder/json.json")
	defer file.Close()

	jsonDecoder := json.NewDecoder(file)

	// jsonStr := `{"Name":"zhoujl", "Age":20}`

	var u1 User
	err := jsonDecoder.Decode(&u1)
	if err != nil {
		log.Printf("jsonEncoder.Encode %v", err)
	}
	
	fmt.Println(u1)
}
