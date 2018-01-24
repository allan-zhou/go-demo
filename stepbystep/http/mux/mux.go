package main

import (
	"fmt"
	"log"
	"net/http"
)

type myMux struct{}

func (mux *myMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		sayHello(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=============URL=============")
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)	
	

	fmt.Fprintf(w, "hello go mux")
}

func main() {
	mux := &myMux{}

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
