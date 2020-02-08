package controller

/*
#cgo CPPFLAGS: -I/home/xiaofo/sgxsdk/include -I/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client
#cgo LDFLAGS: -L/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client -ltee

#include "app.h"
*/
import "C"

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/sslab-instapay/instapay-tee-client/config"
	"github.com/sslab-instapay/instapay-tee-client/service"
	"github.com/sslab-instapay/instapay-tee-client/repository"
	"log"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"strconv"
	"context"
	"fmt"
	"reflect"
	"unsafe"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"encoding/hex"
)

func AccountInformationHandler(context *gin.Context) {
	// TODO Balance 관련 테스트 요망
	account := config.GetAccountConfig()
	balance, _ := service.GetBalance()

	lockedBalance, err := repository.GetAllChannelsLockedBalance()
	if err != nil {
		log.Print(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		convertedBalance, _ := balance.Int64()
		totalBalance := convertedBalance - lockedBalance
		context.JSON(http.StatusOK, gin.H{"address": account.PublicKeyAddress, "balance": totalBalance})
	}

}

//func CreateAccountHandler(context *gin.Context){
//
//	var defaultDirectory string
//	var kf *C.char
//	if os.Getenv("key_file") == ""{
//		defaultDirectory = "data/key/k0"
//		kf = C.CString(defaultDirectory)
//	}else{
//		defaultDirectory = os.Getenv("key_file")
//		kf = C.CString(defaultDirectory)
//	}
//
//	C.ecall_create_account_w()
//	C.ecall_store_account_data_w(kf)
//	defer C.free(unsafe.Pointer(kf))
//
//	var paddrs unsafe.Pointer
//
//	paddrs = C.ecall_get_public_addrs_w()
//	paddrSize := 20
//	paddrSlice := (*[1 << 30]C.address)(unsafe.Pointer(paddrs))[:paddrSize:paddrSize]
//
//	var convertedAddress string
//	convertedAddress = fmt.Sprintf("%02x", paddrSlice[0].addr)
//	convertedAddress = "0x" + convertedAddress
//	fmt.Println("---- Public Key Address ---")
//	fmt.Println(convertedAddress)
//	config.SetAccountConfig(convertedAddress)
//
//	context.JSON(http.StatusOK, gin.H{"address": config.GetAccountConfig().PublicKeyAddress})
//}

func OnchainPaymentHandler(ctx *gin.Context){
	amount, err := strconv.Atoi(ctx.PostForm("amount"))
	otherAddress := ctx.PostForm("address")
	if err != nil{
		log.Println(err)
	}

	client, err := ethclient.Dial("ws://" + config.EthereumConfig["wsHost"] + ":" + config.EthereumConfig["wsPort"])

	fromAddress := common.HexToAddress(config.GetAccountConfig().PublicKeyAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
	}
	convertNonce := C.uint(nonce)
	owner := []C.uchar(config.GetAccountConfig().PublicKeyAddress[2:])
	receiver := []C.uchar(otherAddress[2:])
	convertedAmount := C.uint(amount)
	SigLen := C.uint(0)

	sig := C.ecall_onchain_payment_w(convertNonce, &owner[0], &receiver[0], convertedAmount, &SigLen)
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
	err = client.SendTransaction(context.Background(), tx)
	if err != nil{
		log.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "payment success"})
}