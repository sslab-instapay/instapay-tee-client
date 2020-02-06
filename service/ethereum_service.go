package service

/*
#cgo CPPFLAGS: -I/home/xiaofo/sgxsdk/include -I/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client
#cgo LDFLAGS: -L/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client -ltee

#include "app.h"
*/
import "C"
import (
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	instapay "github.com/sslab-instapay/instapay-tee-client/contracts"
	serverPb "github.com/sslab-instapay/instapay-tee-client/proto/server"
	"github.com/sslab-instapay/instapay-tee-client/config"
	"context"
	"fmt"
	"github.com/sslab-instapay/instapay-tee-client/model"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"math"
	"google.golang.org/grpc"
	"time"
	"reflect"
	"unsafe"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/rlp"
	"os"
)

func SendOpenChannelTransaction(deposit int, otherAddress string) (string, error) {

	client, err := ethclient.Dial("ws://" + config.EthereumConfig["wsHost"] + ":" + config.EthereumConfig["wsPort"])
	if err != nil{
		log.Println(err)
	}

	account := config.GetAccountConfig()
	address := common.HexToAddress(config.GetAccountConfig().PublicKeyAddress)
	nonce, err := client.PendingNonceAt(context.Background(), address)

	convertNonce := C.uint(nonce)
	owner := []C.uchar(account.PublicKeyAddress[2:])
	receiver := []C.uchar(otherAddress[2:])
	newDeposit := C.uint(uint32(deposit))
	SigLen := C.uint(0)

	var sig *C.uchar = C.ecall_create_channel_w(convertNonce, &owner[0], &receiver[0], newDeposit, &SigLen)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(sig)),
		Len:  int(SigLen),
		Cap:  int(SigLen),
	}

	s := *(*[]C.uchar)(unsafe.Pointer(&hdr))
	var convertedRawTx string
	convertedRawTx = fmt.Sprintf("%02x", s)
	rawTxBytes, err := hex.DecodeString(convertedRawTx)
	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)
	client.SendTransaction(context.Background(), tx)

	defer C.free(unsafe.Pointer(sig))

	return "", nil
}

func SendCloseChannelTransaction(channelId int64) {

	client, err := ethclient.Dial("ws://" + config.EthereumConfig["wsHost"] + ":" + config.EthereumConfig["wsPort"])
	if err != nil {
		log.Println(err)
	}
	address := common.HexToAddress(config.GetAccountConfig().PublicKeyAddress)
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Println(err)
	}

	ChannelID := C.uint(channelId)
	SigLen := C.uint(0)

	var sig2 *C.uchar = C.ecall_close_channel_w(C.uint(nonce), ChannelID, &SigLen)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(sig2)),
		Len:  int(SigLen),
		Cap:  int(SigLen),
	}

	s := *(*[]C.uchar)(unsafe.Pointer(&hdr))
	var convertedRawTx string
	convertedRawTx = fmt.Sprintf("%02x", s)
	rawTxBytes, err := hex.DecodeString(convertedRawTx)
	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)
	client.SendTransaction(context.Background(), tx)

}

func ListenContractEvent() {
	log.Println("---Start Listen Contract Event---")
	client, err := ethclient.Dial("ws://" + config.EthereumConfig["wsHost"] + ":" + config.EthereumConfig["wsPort"])
	if err != nil {
		log.Fatal("Cannot connect Ethereum So, End Client")
	}
	contractAddress := common.HexToAddress(config.EthereumConfig["contractAddr"])

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Println(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(instapay.InstapayABI)))
	if err != nil {
		log.Println(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Println(err)
		case vLog := <-logs:
			var createChannelEvent = model.CreateChannelEvent{}
			var closeChannelEvent = model.CloseChannelEvent{}
			var ejectEvent = model.EjectEvent{}

			err := contractAbi.Unpack(&createChannelEvent, "EventCreateChannel", vLog.Data)
			if err == nil {
				log.Println("CreateChannel Event Emission")
				fmt.Printf("Channel ID       : %d\n", createChannelEvent.Id)
				fmt.Printf("Channel Onwer    : %s\n", createChannelEvent.Owner.Hex())
				fmt.Printf("Channel Receiver : %s\n", createChannelEvent.Receiver.Hex())
				fmt.Printf("Channel Deposit  : %d\n", createChannelEvent.Deposit)
				HandleCreateChannelEvent(createChannelEvent)
			}

			err = contractAbi.Unpack(&closeChannelEvent, "EventCloseChannel", vLog.Data)
			if err == nil {
				log.Print("CloseChannel Event Emission")
				fmt.Printf("Channel ID       : %d\n", closeChannelEvent.Id)
				fmt.Printf("Owner Balance    : %d\n", closeChannelEvent.Ownerbal)
				fmt.Printf("Receiver Balance : %d\n", closeChannelEvent.Receiverbal)
				HandleCloseChannelEvent(closeChannelEvent)
			}

			err = contractAbi.Unpack(&ejectEvent, "EventEject", vLog.Data)
			if err == nil {
				fmt.Printf("Payment Number   : %d\n", ejectEvent.Pn)
				fmt.Printf("Stage            : %d\n", ejectEvent.Registeredstage)
			}

		}
	}
}

func HandleCreateChannelEvent(event model.CreateChannelEvent) error{

	account := config.GetAccountConfig()
	log.Println("----- Handle Create Channel Event ----")

	if event.Receiver.String() == config.GetAccountConfig().PublicKeyAddress {
		// CASE IN CHANNEL
		channelId := C.uint(uint32(event.Id))
		owner := []C.uchar(account.PublicKeyAddress[2:])
		sender := []C.uchar(event.Receiver.Hex()[2:])
		deposit := C.uint(uint32(event.Deposit))
		C.ecall_receive_create_channel_w(channelId, &sender[0], &owner[0], deposit)
	} else if event.Owner.String() == config.GetAccountConfig().PublicKeyAddress {
		// CASE OUT CHANNEL
		channelId := C.uint(uint32(event.Id))
		owner := []C.uchar(event.Receiver.Hex()[2:])
		sender := []C.uchar(account.PublicKeyAddress[2:])
		deposit := C.uint(uint32(event.Deposit))
		C.ecall_receive_create_channel_w(channelId, &sender[0], &owner[0], deposit)
	}

	connection, err := grpc.Dial(config.EthereumConfig["serverGrpcHost"] + ":" + config.EthereumConfig["serverGrpcPort"], grpc.WithInsecure())
	if err != nil {
		log.Println("GRPC Connection Error")
		log.Println(err)
		return err
	}
	defer connection.Close()
	client := serverPb.NewServerClient(connection)

	clientContext, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var otherAddress string
	if event.Receiver.String() == config.GetAccountConfig().PublicKeyAddress{
		otherAddress = event.Owner.String()
	}else {
		otherAddress = event.Receiver.String()
	}

	// TODO 상대 아이피 주소  이런 것들 업데이트
	_, err = client.CommunicationInfoRequest(clientContext, &serverPb.Address{Addr: otherAddress})
	if err != nil {
		log.Println(err)
		return err
	}

	var defaultDirectory string
	if os.Getenv("channel_file") == "" {
		defaultDirectory = "./data/channel/c0"
	}else{
		defaultDirectory = os.Getenv("channel_file")
	}

	cf := C.CString(defaultDirectory)
	C.ecall_store_channel_data_w(cf)

	log.Println("----- Handle Create Channel Event END ----")
	return nil
}

func HandleCloseChannelEvent(event model.CloseChannelEvent) {

	log.Println("----- Handle Close Channel Event -----")
	channelId := C.uint(uint32(event.Id))
	ownerBal := C.uint(uint32(event.Ownerbal))
	receiverBal := C.uint(uint32(event.Receiverbal))

	log.Println("----- Start Close Channel Event -----")
	C.ecall_receive_close_channel_w(channelId, ownerBal, receiverBal)

	var defaultDirectory string
	if os.Getenv("channel_file") == "" {
		defaultDirectory = "./data/channel/c0"
	}else{
		defaultDirectory = os.Getenv("channel_file")
	}

	cf := C.CString(defaultDirectory)
	C.ecall_store_channel_data_w(cf)
	log.Println("----- End Close Channel Event -----")

}

func HandleEjectEvent(event model.EjectEvent) {
	//TODO
}

func GetBalance() (big.Float, error) {

	account := common.HexToAddress(config.GetAccountConfig().PublicKeyAddress)
	client, err := ethclient.Dial("ws://" + config.EthereumConfig["wsHost"] + ":" + config.EthereumConfig["wsPort"])

	if err != nil {
		return *big.NewFloat(0), err
	}

	balance, err := client.BalanceAt(context.Background(), account, nil)

	if err != nil {
		return *big.NewFloat(0), err
	}
	log.Println(balance)

	floatBalance := new(big.Float)
	floatBalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(floatBalance, big.NewFloat(math.Pow10(18)))

	return *ethValue, nil
}
