package main

import (
	"crypto/x509"
	// "crypto/rsa"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func main() {
	filepath := "../../certs/rsa/server.crt"
	fileBytes, _ := ioutil.ReadFile(filepath)
	// fmt.Println(fileBytes)
	// fmt.Println(string(fileBytes))

	block, _ := pem.Decode(fileBytes)
	// fmt.Println(rest)
	// fmt.Printf("block Type:\n %s \n", block.Type)
	// fmt.Printf("block Bytes:\n %v \n", block.Bytes)
	// fmt.Printf("block Headers:\n %v \n", block.Headers)

	cert, _ := x509.ParseCertificate((block.Bytes))
	fmt.Printf("Version:\n %v \n", cert.Version)
	fmt.Printf("Subject:\n %v \n", cert.Subject)
	fmt.Printf("PublicKeyAlgorithm:\n %v \n", cert.PublicKeyAlgorithm)
	fmt.Printf("PublicKey:\n %v \n", cert.PublicKey)
	// fmt.Printf("RawSubjectPublicKeyInfo:\n %v \n", cert.RawSubjectPublicKeyInfo)

	fmt.Printf("PublicKey:\n %v \n", &cert.PublicKey)
}
