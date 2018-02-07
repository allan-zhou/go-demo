package main

import (
	"fmt"
	// "crypto/elliptic"
	"encoding/asn1"	
)
type secgNamedCurve asn1.ObjectIdentifier

var (
	secgNamedCurveP224 = secgNamedCurve{1, 3, 132, 0, 33}
)

func main()  {		
	fmt.Println(*&secgNamedCurveP224)	

	b := []byte{6, 5, 4, 3, 1, 2, 9, 4, 0, 3, 3}
	fmt.Println(string(b))
}