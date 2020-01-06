package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	instapayContract "github.com/sslab-instapay/instapay-tee-client/contracts"
	"math/big"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"context"
)

func main(){
	client, err := ethclient.Dial("ws://141.223.121.139:8881")
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("e113ff405699b7779fbe278ee237f2988b1e6769d586d8803860d49f28359fbd")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	contract, _ := instapayContract.NewInstapay(fromAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(contract)
}
