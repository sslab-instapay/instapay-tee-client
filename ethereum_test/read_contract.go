package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/common"
	store "github.com/sslab-instapay/instapay-tee-client/contract"
	"math/big"
	"fmt"
)

func main(){
	client, err := ethclient.Dial("ws://141.223.121.139:8881")
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0x092d70BB5c1954F5Fa3EBbb282d0416a5e46c818")
	instance, err := store.NewContract(address, client)
	if err != nil {
		log.Fatal(err)
	}
	channels, err := instance.Channels(nil, big.NewInt(20))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(channels)
}
