package aes

import (
	"crypto/cipher"
	"crypto/aes"
	"crypto/ecdsa"
	"crypto/x509"
	"crypto/rand"
	"io/ioutil"
	"io"
	"encoding/pem"
	"fmt"
	"encoding/base64"
)

type ServerPems struct {
	privateKey *ecdsa.PrivateKey
	cert       *x509.Certificate
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


func Encrypt(secret []byte, plaintext []byte) ([]byte) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(block)

	cipherText := make([]byte, len(plaintext) + aes.BlockSize)	
	iv := cipherText[:aes.BlockSize]
	io.ReadFull(rand.Reader, iv)
	fmt.Println(iv)

	stream := cipher.NewCTR(block, iv)	
	stream.XORKeyStream(cipherText[aes.BlockSize:], plaintext)

	return cipherText
}

func Decrypt(secret []byte, ciphertext []byte) ([]byte) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		fmt.Println(err)
	}

	plaintext := make([]byte, len(ciphertext) - aes.BlockSize)
	iv := ciphertext[:aes.BlockSize]

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, ciphertext[aes.BlockSize:])
	
	return plaintext
}

func main()  {
	secret := make([]byte,16)
	secretString := "mykey"
	copy(secret[:], secretString)
	fmt.Println(secret)

	plaintext := []byte("hello crypto")
	fmt.Println(plaintext)

	//Encrypt
	ciphertext := Encrypt(secret, plaintext)
	fmt.Println(ciphertext)
	fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))

	//Decrypt
	plaintext = Decrypt(secret, ciphertext)
	fmt.Println(plaintext)
	fmt.Println(string(plaintext))
}