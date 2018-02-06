package main

import (
	"fmt"
	"strings"	
	"hash"
	"crypto/md5"
	"crypto"
	"encoding/base64"
	"encoding/hex"
)

type Person struct {
	firstName string
	lastName  string	
	hash func() hash.Hash
	hashAlgo crypto.Hash
}

func upperPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	var person11 Person
	person11.firstName = "jialiang"
	person11.lastName = "zhou"
	person11.hash = md5.New
	person11.hashAlgo = crypto.SHA256
	// person11.hash.Sum([]byte("hello"))

	h := person11.hash().Sum([]byte("hello hash"))
	fmt.Println(h)
	fmt.Println(hex.EncodeToString(h))
	fmt.Println(base64.StdEncoding.EncodeToString(h))
	fmt.Println(person11)
	fmt.Printf("%s\n %s\n %v\n %v\n ", person11.firstName, person11.lastName, person11.hash, person11.hashAlgo.HashFunc())
}
