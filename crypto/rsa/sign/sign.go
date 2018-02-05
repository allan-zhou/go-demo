package main

import (
	"crypto"
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
	privateKey *rsa.PrivateKey
	cert       *x509.Certificate
}

func NewServerPems(keyFileName string, certFileName string) (*ServerPems, error) {
	key, _ := loadPrivateKey(keyFileName)

	cert, _ := loadX509Cert(certFileName)

	return &ServerPems{key, cert}, nil
}

func loadPrivateKey(keyFileName string) (*rsa.PrivateKey, error) {
	file, _ := ioutil.ReadFile(keyFileName)

	block, _ := pem.Decode(file)

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func loadX509Cert(certFileName string) (*x509.Certificate, error) {
	file, _ := ioutil.ReadFile(certFileName)

	block, _ := pem.Decode(file)

	return x509.ParseCertificate(block.Bytes)
}

func (pem *ServerPems) Sign(data []byte) ([]byte, error) {
	hashed := sha256.Sum256(data)
	return rsa.SignPKCS1v15(rand.Reader, pem.privateKey, crypto.SHA256, hashed[:])
}

func (pem *ServerPems) Verify(data []byte, sig []byte) error {
	hashed := sha256.Sum256(data)
	return rsa.VerifyPKCS1v15(pem.cert.PublicKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], sig)
}

func main() {
	message := []byte("hello")

	pem, _ := NewServerPems("../../certs/rsa/server.key", "../../certs/rsa/server.crt")

	sig, _ := pem.Sign(message)
	fmt.Println(base64.StdEncoding.EncodeToString(sig))

	err := pem.Verify(message, sig)
	fmt.Println(err)
}
