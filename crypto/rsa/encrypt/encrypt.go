package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

type ServerPems struct {
	privatekey *rsa.PrivateKey
	cert       *x509.Certificate
}

func NewServerPems(keyFile string, certFile string) (*ServerPems, error) {
	privatekey, _ := loadPrivateKey(keyFile)

	cert, _ := loadX509Cert(certFile)

	return &ServerPems{privatekey, cert}, nil
}

func loadPrivateKey(keyFilePath string) (*rsa.PrivateKey, error) {
	file, _ := ioutil.ReadFile(keyFilePath)

	block, _ := pem.Decode(file)

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func loadX509Cert(certFilePath string) (*x509.Certificate, error) {
	file, _ := ioutil.ReadFile(certFilePath)

	block, _ := pem.Decode(file)

	return x509.ParseCertificate(block.Bytes)
}

func (pem *ServerPems) Encrypt(message []byte, label []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pem.cert.PublicKey.(*rsa.PublicKey), message, label)
}

func (pem *ServerPems) Decrypt(data []byte, label []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, pem.privatekey, data, label)
}

func main() {
	pems, _ := NewServerPems("../../certs/rsa/server.key", "../../certs/rsa/server.crt")

	label := []byte("hello")
	message := []byte("hello")

	enData, _ := pems.Encrypt(message, label)

	fmt.Println(base64.StdEncoding.EncodeToString(enData))

	deData, _ := pems.Decrypt(enData, label)

	fmt.Println(string(deData))
}
