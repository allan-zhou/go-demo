package main

import (
	"fmt"
	// "hash"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func loadPrivateKey()  {
	
}

func Encrypt()  {
	
}

func main()  {	
	h := sha256.New()
	
	h.Write([]byte("hello"))
	bs := h.Sum(nil)	

	fmt.Printf("base64 encoding:\n %v \n", bs)
	fmt.Printf("base64 encoding:\n %s \n", base64.StdEncoding.EncodeToString(bs))
	fmt.Printf("base64 encoding:\n %s \n", hex.EncodeToString(bs))

	fmt.Printf("sha256.Sum256: \n %v \n",sha256.Sum256([]byte("hello")))
}