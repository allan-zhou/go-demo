package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"encoding/asn1"
	"encoding/base64"
)

type ServerPems struct {
	privateKey *ecdsa.PrivateKey
	cert       *x509.Certificate
}

type ECDSASignature struct {
	R,S *big.Int
}

// 04:
// 34:f4:c8:cf:30:34:64:8b:cc:ed:49:69:30:63:67:8e:
// 40:9e:df:06:ce:33:50:f8:65:4d:db:5a:ca:c6:b3:1e:
// 5c:c5:5a:64:4f:e6:d3:a5:ad:26:c9:df:fa:f2:45:5e:
// f4:78:c8:f1:af:39:96:ff:b7:ad:75:73:69:b7:1c:99

func  NewServerPems(keyFileName string, certFileName string) (*ServerPems, error) {
	key,_ := loadPrivatekey(keyFileName)
	// fmt.Println(key)

	cert, _ := loadX509Cert(certFileName)
	fmt.Println(cert.SignatureAlgorithm)
	fmt.Printf("%X \n", cert.Signature)
	fmt.Println(cert.PublicKeyAlgorithm)
	pk := cert.PublicKey.(*ecdsa.PublicKey)
	fmt.Println(pk)
	fmt.Printf("%X \n",pk.X)
	fmt.Printf("%X \n",pk.Y)

	return &ServerPems{key, cert}, nil
}

func loadPrivatekey(keyFileName string) (*ecdsa.PrivateKey, error) {
	file, _ := ioutil.ReadFile(keyFileName)

	block, _ := pem.Decode(file)

	return x509.ParseECPrivateKey(block.Bytes)
}

func loadX509Cert(certFileName string) (*x509.Certificate, error) {
	file, _ := ioutil.ReadFile(certFileName)

	block, _ := pem.Decode(file)

	return x509.ParseCertificate(block.Bytes)
}

func (pem *ServerPems) Sign(data []byte) ([]byte, error) {
	hashed := sha256.Sum256(data)
	r,s,_ := ecdsa.Sign(rand.Reader, pem.privateKey, hashed[:])

	raw, _ := asn1.Marshal(ECDSASignature{r, s})	

	return raw, nil
}

func (pem *ServerPems) Verify(data []byte,sig []byte) (bool) {
	hashed := sha256.Sum256(data)

	ecdsaSignature := new(ECDSASignature)
	_, err := asn1.Unmarshal(sig, ecdsaSignature)
	
	if err != nil {
		fmt.Println(err)
	}

	return ecdsa.Verify(pem.cert.PublicKey.(*ecdsa.PublicKey), hashed[:], ecdsaSignature.R, ecdsaSignature.S)
}

func main() {
	message := []byte("hello")

	pem,_ := NewServerPems("../../certs/ecc/prime256v1/server.key", "../../certs/ecc/prime256v1/server.crt")

	sig ,_ := pem.Sign(message)
	fmt.Println(base64.StdEncoding.EncodeToString(sig))

	ret := pem.Verify(message,sig)
	fmt.Println(ret)

	// 23952764745190639315453386127001698040085062535617457724687902786301481890590 
	// 41961474803002354966835355886584688754364104001902783508288901242804267064473
}
