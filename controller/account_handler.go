package controller

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