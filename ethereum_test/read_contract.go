package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	instapayContract "github.com/sslab-instapay/instapay-tee-server/contracts"
	"log"
)

func main(){
	client, err := ethclient.Dial("ws://141.223.121.139:8881")
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0x092d70BB5c1954F5Fa3EBbb282d0416a5e46c818")
	instance, err := instapayContract.NewInstapay(address , client)
	if err != nil {
		log.Fatal(err)
	}

	readme, err := instance.Readme(nil)
	if err != nil {
		fmt.Printf("ERROR 2: ")
		log.Fatal(err)
	}

	fmt.Println(readme)

	deposit, err := instance.Getchanneldeposit(nil, big.NewInt(5))
	if err != nil {
		fmt.Printf("ERROR 2: ")
		log.Fatal(err)
	}

	fmt.Println(deposit)

	status, err := instance.Getchannelstatus(nil, big.NewInt(5))
	if err != nil {
		fmt.Printf("ERROR 2: ")
		log.Fatal(err)
	}

	fmt.Println(status)
}