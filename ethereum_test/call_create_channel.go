package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	instapay "github.com/sslab-instapay/instapay-tee-client/contracts"
	"github.com/sslab-instapay/instapay-tee-client/config"
)

func main() {
	client, err := ethclient.Dial("ws://141.223.121.139:8881")
	if err != nil {
		log.Fatal(err)
	}

	// loading instapay contract on the blockchain
	address := common.HexToAddress(config.EthereumConfig["contractAddr"])  // change to correct address
	instance, err := instapay.NewInstapay(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// loading my public key, nonce and gas price
	privateKey, err := crypto.HexToECDSA(config.GetAccountConfig().PrivateKey)
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

	// composing a transaction
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(80000000000) // in wei
	auth.GasLimit = uint64(2000000) // in units
	auth.GasPrice = gasPrice

	receiver := common.HexToAddress("0x0b4161ad4f49781a821C308D672E6c669139843C")

	tx, err := instance.CreateChannel(auth, receiver)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}
