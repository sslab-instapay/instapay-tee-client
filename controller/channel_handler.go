package controller

/*
#cgo CPPFLAGS: -I/home/xiaofo/sgxsdk/include -I/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client
#cgo LDFLAGS: -L/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client -ltee

#include "app.h"
*/
import "C"

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/sslab-instapay/instapay-tee-client/config"
	"github.com/sslab-instapay/instapay-tee-client/model"
	clientPb "github.com/sslab-instapay/instapay-tee-client/proto/client"
	serverPb "github.com/sslab-instapay/instapay-tee-client/proto/server"
	"github.com/sslab-instapay/instapay-tee-client/repository"
	"github.com/sslab-instapay/instapay-tee-client/service"
	"github.com/sslab-instapay/instapay-tee-client/util"
	"google.golang.org/grpc"
)

var ExecutionTime time.Time

func OpenChannelHandler(ctx *gin.Context) {

	otherAddress := ctx.PostForm("other_addr")
	deposit, _ := strconv.Atoi(ctx.PostForm("deposit"))
	fmt.Println(otherAddress)

	txHash, err := service.SendOpenChannelTransaction(deposit, otherAddress)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"result": "success", "txHash": txHash})
	}
}

func CloseChannelHandler(ctx *gin.Context) {
	channelIdParam := ctx.PostForm("channelId")
	log.Println(channelIdParam)
	channelId, _ := strconv.Atoi(channelIdParam)
	log.Println(channelId)

	service.SendCloseChannelTransaction(int64(channelId))

	ctx.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func EjectChannelHandler(ctx *gin.Context) {

}

// TODO 데모 시나리오 이후 구현
func DepositChannelHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func DirectPayChannelHandler(ctx *gin.Context) {
	channelIdParam := ctx.PostForm("ch_id")
	amountParam := ctx.PostForm("amount")

	channelId, err := strconv.Atoi(channelIdParam)
	if err != nil {
		log.Println(err)
	}

	amount, err := strconv.Atoi(amountParam)
	if err != nil {
		log.Println(err)
	}

	channel, err := repository.GetChannelById(channelId)
	if err != nil {
		log.Println(err)
	}

	peerInformation, _, err := util.GetPeerInformationByAddress(channel.OtherAddress)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	conn, err := grpc.Dial(peerInformation.IpAddress+":"+strconv.Itoa(peerInformation.GrpcPort), grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect: %v", err)
	}
	// pay_w 실행 후 상대에게 요청
	var originalMessage *C.uchar
	var signature *C.uchar

	C.ecall_pay_w(C.uint(uint32(channelId)), C.uint(uint32(amount)), &originalMessage, &signature)
	defer conn.Close()
	client := clientPb.NewClientClient(conn)
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	originalMessageByte, signatureByte := convertPointerToByte(originalMessage, signature)
	r, err := client.DirectChannelPayment(ctx, &clientPb.DirectChannelPaymentMessage{ChannelId: int64(channelId), Amount: int64(amount), OriginalMessage: originalMessageByte, Signature: signatureByte})
	if err != nil {
		log.Println("could not greet: %v", err)
	}

	log.Println(r.Result)
	if r.Result {
		originalMessagePointer, signaturePointer := convertByteToPointer(r.ReplyMessage, r.ReplySignature)
		C.ecall_pay_accepted_w(originalMessagePointer, signaturePointer)
		log.Println("----- payment accept w end -----")
		ctx.JSON(http.StatusOK, gin.H{"success": r.Result})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": r.Result})
	}
}

func PaymentToServerChannelHandler(ctx *gin.Context) {

	otherAddress := ctx.PostForm("addr")
	amount, err := strconv.Atoi(ctx.PostForm("amount"))
	if err != nil {
		log.Println(err)
	}

	myAddress := config.GetAccountConfig().PublicKeyAddress
	connection, err := grpc.Dial(config.EthereumConfig["serverGrpcHost"]+":"+config.EthereumConfig["serverGrpcPort"], grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	defer connection.Close()
	client := serverPb.NewServerClient(connection)

	clientContext, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ExecutionTime = time.Now()
	r, err := client.PaymentRequest(clientContext, &serverPb.PaymentRequestMessage{From: myAddress, To: otherAddress, Amount: int64(amount)})
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	log.Println(r.GetResult())

	ctx.JSON(http.StatusOK, gin.H{"sendAddress": otherAddress, "amount": amount})
}

func GetChannelListHandler(ctx *gin.Context) {

	channelList, err := repository.GetOpenedChannelList()
	if err != nil {
		log.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"channels": channelList,
	})
}

func GetWalletInformationHandler(ctx *gin.Context) {

	account := config.GetAccountConfig()

	tempBalance, err := service.GetBalance()
	var balance int64
	if err != nil {
		log.Println(err)
	}

	balance, _ = tempBalance.Int64()

	offchainDeposit, err := repository.GetAllDepositValue()
	if err != nil {
		log.Println(err)
	}

	offchainBalance, err := repository.GetOffChainBalance()
	if err != nil {
		log.Println(err)
	}

	accountDto := model.AccountDTO{
		PublicKeyAddress: account.PublicKeyAddress,
		Balance:          int(balance),
		OffChainDeposit:  offchainDeposit,
		OffChainBalance:  offchainBalance,
	}

	openedChannelList, err := repository.GetOpenedChannelList()
	if err != nil {
		log.Println(err)
	}

	inChannelList := make([]model.Channel, 0)
	outChannelList := make([]model.Channel, 0)

	for _, channel := range openedChannelList {
		if channel.Type == model.IN {
			inChannelList = append(inChannelList, channel)
		} else {
			outChannelList = append(outChannelList, channel)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"account": accountDto, "inChannelList": inChannelList, "outChannelList": outChannelList,
	})

}

func convertByteToPointer(originalMsg []byte, signature []byte) (*C.uchar, *C.uchar) {

	var uOriginal [44]C.uchar
	var uSignature [65]C.uchar

	for i := 0; i < 44; i++ {
		uOriginal[i] = C.uchar(originalMsg[i])
	}

	for i := 0; i < 65; i++ {
		uSignature[i] = C.uchar(signature[i])
	}

	cOriginalMsg := (*C.uchar)(unsafe.Pointer(&uOriginal[0]))
	cSignature := (*C.uchar)(unsafe.Pointer(&uSignature[0]))

	return cOriginalMsg, cSignature
}

func convertPointerToByte(originalMsg *C.uchar, signature *C.uchar) ([]byte, []byte) {

	var returnMsg []byte
	var returnSignature []byte

	replyMsgHdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(originalMsg)),
		Len:  int(44),
		Cap:  int(44),
	}
	replyMsgS := *(*[]C.uchar)(unsafe.Pointer(&replyMsgHdr))

	replySigHdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(signature)),
		Len:  int(65),
		Cap:  int(65),
	}
	replySigS := *(*[]C.uchar)(unsafe.Pointer(&replySigHdr))

	for i := 0; i < 44; i++ {
		returnMsg = append(returnMsg, byte(replyMsgS[i]))
	}

	for i := 0; i < 65; i++ {
		returnSignature = append(returnSignature, byte(replySigS[i]))
	}

	return returnMsg, returnSignature
}
