package main

import (
	"context"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	instapay "github.com/sslab-instapay/instapay-go-client/contract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"github.com/sslab-instapay/instapay-go-client/model"
	"github.com/sslab-instapay/instapay-go-client/repository"
	"github.com/sslab-instapay/instapay-go-client/config"
	)

var EthereumConfig = map[string]string{
	/* web3 and ethereum */
	"rpcHost":          "141.223.121.139",
	"rpcPort":          "8555",
	"wsHost":           "141.223.121.139", //141.223.121.139
	"wsPort":           "8881",
	"contractAddr":     "0x3016947BE73dcb877401Ee33802aC8fA6feE631E", // change to correct address
	"contractSrcPath":  "../contracts/InstaPay.sol",
	"contractInstance": "",
	"web3":             "",
	"event":            "",

	/* grpc configuration */
	"serverGrpcHost": "141.223.121.139",
	"serverGrpcPort": "50004",
	"serverProto":    "",
	"server":         "",
	"myGrpcPort":     "", //process.argv[3]
	"clientProto":    "",
	"receiver":       "",
}

func main() {
	client, err := ethclient.Dial("ws://" + config.EthereumConfig["wsHost"] + ":" + config.EthereumConfig["wsPort"])
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(config.EthereumConfig["contractAddr"])

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(instapay.ContractABI)))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			var createChannelEvent = model.CreateChannelEvent{}
			var closeChannelEvent = model.CloseChannelEvent{}
			var ejectEvent = model.EjectEvent{}

			log.Println("---Event Receive---")
			err := contractAbi.Unpack(&createChannelEvent, "EventCreateChannel", vLog.Data)
			if err == nil {
				HandleCreateChannelEvent(createChannelEvent)
			}

			err = contractAbi.Unpack(&closeChannelEvent, "EventCloseChannel", vLog.Data)
			if err == nil {
				log.Println("CloseChannel Event Emission")
				HandleCloseChannelEvent(closeChannelEvent)
			}

			err = contractAbi.Unpack(&ejectEvent, "EventEject", vLog.Data)
			if err == nil {

			}
			log.Println("---Event Handling END---")

		}
	}
}

func HandleCreateChannelEvent(event model.CreateChannelEvent) {

	var channel = model.Channel{ChannelId: event.Id.Int64(),
		Status: model.IDLE, MyAddress: event.Receiver.String(),
		MyBalance: 0, MyDeposit: 0, OtherAddress: event.Owner.String()}
	// TODO server에 보냄

	repository.InsertChannel(channel)
}

func HandleCloseChannelEvent(event model.CloseChannelEvent) {
	channel, err := repository.GetChannelById(event.Id.Int64())

	if err != nil {
		log.Println("there is no channel")
		return
	}

	channel.Status = model.CLOSED
	repository.UpdateChannel(channel)
}
