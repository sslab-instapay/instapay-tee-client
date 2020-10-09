package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	instapay "github.com/sslab-instapay/instapay-tee-client/contracts"
)

func main() {
	client, err := ethclient.Dial("ws://141.223.121.139:8881")
	if err != nil {
		log.Fatal(err)
	}

	contract, err := instapay.NewInstapay(common.HexToAddress("0x745a8d1610D4AC940350221F569338E4C93b1De6"), client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(contract.Readme(nil))
	channel, err := contract.Channels(nil, big.NewInt(1))
	fmt.Println(channel.Deposit)
	fmt.Println(channel.Status)
	fmt.Println(channel.Owner.Hex())
	fmt.Println(channel.Receiver.Hex())
}
