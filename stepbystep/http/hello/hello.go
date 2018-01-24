package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {	
	fmt.Println("Method",r.Method)

	fmt.Println("=============URL=============")
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)	
	
	fmt.Println("=============Proto=============")
	fmt.Println("Proto",r.Proto)
	fmt.Println("ProtoMajor",r.ProtoMajor)
	fmt.Println("ProtoMinor",r.ProtoMinor)

	fmt.Println("=============Header=============")
	fmt.Println("Header",r.Header)

	fmt.Println("=============body=============")
	fmt.Println("Body",r.Body)

	fmt.Fprintf(w, "hello go")
}
func main() {
	http.HandleFunc("/hello", helloHandle)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
