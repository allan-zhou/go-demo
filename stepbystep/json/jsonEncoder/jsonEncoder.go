package main

import (
	"encoding/json"
	"log"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	file, _ := os.OpenFile("json.json", os.O_CREATE|os.O_WRONLY, 0744)
	defer file.Close()

	jsonEncoder := json.NewEncoder(file)

	// jsonStr := `{"Name":"zhoujl", "Age":20}`

	u1 := &User{"zhoujl", 20}
	err := jsonEncoder.Encode(u1)
	if err != nil {
		log.Printf("jsonEncoder.Encode ", err)
	}
	
}
