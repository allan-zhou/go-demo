package main

import (
	"io"
	"fmt"
	"crypto/ecdsa"
	"crypto/x509"
	"crypto/aes"
	"crypto/cipher"
	"crypto/elliptic"
	"crypto/rand"
	"io/ioutil"
	"encoding/pem"
	// "encoding/base64"
	"math/big"
)

type ServerPems struct {
	privateKey *ecdsa.PrivateKey
	cert *x509.Certificate
}

type PrivateKey struct {
	PublicKey
	D *big.Int
}

type PublicKey struct {
	elliptic.Curve
	X, Y *big.Int
}


func NewServerPems(privateKeyFileName string, certFileName string) (*ServerPems, error) {
	key,_ := loadPrivateKey(privateKeyFileName)

	cert,_ := loadX509Cert(certFileName)

	return &ServerPems{key, cert},nil
}

func loadPrivateKey(keyFileName string) (*ecdsa.PrivateKey, error) {
	file, _ := ioutil.ReadFile(keyFileName)

	block, _ := pem.Decode(file)

	return x509.ParseECPrivateKey(block.Bytes)
}

func loadX509Cert(keyFileName string) (*x509.Certificate, error) {
	file,_ := ioutil.ReadFile(keyFileName)

	block,_ := pem.Decode(file)

	return x509.ParseCertificate(block.Bytes)
}


func (privateKey *PrivateKey) GenerateSharedSecret(pubKey *PublicKey) () {
	x,y := privateKey.ScalarMult(pubKey.X, pubKey.Y, privateKey.D.Bytes())
	fmt.Println(x)
	fmt.Println(y)

	return
}

func aesEncrypt(secret []byte, plaintext []byte) (ciphertext []byte) {
	block,_ := aes.NewCipher(secret)

	ciphertext = make([]byte, len(plaintext) + aes.BlockSize)
	iv := ciphertext[:aes.BlockSize]
	io.ReadFull(rand.Reader, iv)

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return
}

func aesDecrypt(secret []byte, ciphertext []byte) (plaintext []byte) {
	block,_ := aes.NewCipher(secret)

	plaintext = make([]byte, len(ciphertext) - aes.BlockSize)
	iv := ciphertext[:aes.BlockSize]

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, ciphertext[aes.BlockSize:])
	return
}

func generateSecret()  {
	
}

func (pem *ServerPems) Encrypt(plaintext []byte) (ciphertext []byte) {
	ciphertext = make([]byte, len(plaintext))

	pubKey := pem.cert.PublicKey.(*ecdsa.PublicKey)

	rb := elliptic.Marshal(elliptic.P256(), pubKey.X, pubKey.Y)
	fmt.Printf("%X \n", rb)

	return
}

func main()  {
	pem,_ := NewServerPems("../../certs/ecc/prime256v1/server.key", "../../certs/ecc/prime256v1/server.crt")
	secret := make([]byte,16)
	secretString := "mykey"
	copy(secret[:], secretString)

	pem.Encrypt([]byte("hello"))
	// ciphertext := aesEncrypt(secret, []byte("hello crypto"))
	// fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))

	// plaintext := aesDecrypt(secret, ciphertext)
	// fmt.Println(string(plaintext))
}
