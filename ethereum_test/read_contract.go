package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	instapayContract "github.com/sslab-instapay/instapay-tee-client/contracts"
	"log"
)

func main(){
	client, err := ethclient.Dial("ws://141.223.121.139:8881")
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0xe8ff3ac1ca790a4656a34a0a82442d15ea24b3f7aa75bfa6f15d003804bb306a")
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
}
