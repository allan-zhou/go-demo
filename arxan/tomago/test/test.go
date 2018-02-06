package main

import (
	"fmt"
	"net/http"
	"github.com/arxanchain/sdk-go-common/structs"
	restapi "github.com/arxanchain/sdk-go-common/rest/api"
	tomagoapi "go-demo/arxan/tomago/api"
)

func main()  {
	
	cryptoConfig := &restapi.CryptoConfig{
		Enable:         true,
		CertsStorePath: "../../certs",
	}
	
	config := &restapi.Config{
		Address:   "http://192.168.1.185:9143",
		ApiKey:    "F8pxCN1Mp1516701344",
		CryptoCfg: cryptoConfig,
	}
	
	tomagoClient, err := tomagoapi.NewTomagoClient(config)
	if err != nil {
		fmt.Println(err)
	}

	chaincodeClient := tomagoClient.GetBlockchainClient()

	chaincodeID := "transfer2"
	
	reqBody := &structs.PayloadWithTags{
		Payload: &structs.ChaincodeRequest{
			ChaincodeID: chaincodeID,
			Args:        []string{"query", "a"},
		},
	}
	
	//set http header
	header := http.Header{}
	header.Set("Channel-Id", "pubchain")
	
	//do query blockchain
	resp, err := chaincodeClient.Query(header, reqBody)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}