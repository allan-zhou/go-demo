package main

import (
	"fmt"
	"github.com/arxanchain/sdk-go-common/crypto/ecc"
	"encoding/base64"
)

func main()  {
	keyfile := "../../certs/ecc/prime256v1/server.key"
	certfile := "../../certs/ecc/prime256v1/server.crt"
	eccLib,_ := ecc.NewECCCryptoLib(keyfile, certfile)
	
	plaintext := []byte("hello")

	ciphertext,_ := eccLib.Encrypt(plaintext)
	fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))

	detext,_ := eccLib.Decrypt(ciphertext)
	fmt.Println(string(detext))
}