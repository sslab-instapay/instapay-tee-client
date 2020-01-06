package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	instapayContract "github.com/sslab-instapay/instapay-tee-client/contracts"
	"log"
)

func main() {
	client, err := ethclient.Dial("ws://141.223.121.139:8881")
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0xD8B5088356F0F6B813a97E0A8DEd3Bb4f7a7e8d6")
	instance, err := instapayContract.NewInstapay(address, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract is loaded")
	_ = instance
}
