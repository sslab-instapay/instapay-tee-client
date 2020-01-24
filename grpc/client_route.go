package grpc

/*
#cgo CPPFLAGS: -I/home/xiaofo/sgxsdk/include -I/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client
#cgo LDFLAGS: -L/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client -ltee

#include "app.h"
*/
import "C"

import (
	"context"
	clientPb "github.com/sslab-instapay/instapay-tee-client/proto/client"
	"github.com/sslab-instapay/instapay-tee-client/controller"
	"log"
	"fmt"
	"time"
			"github.com/sslab-instapay/instapay-tee-client/util"
)

type ClientGrpc struct {
}

func (s *ClientGrpc) AgreementRequest(ctx context.Context, in *clientPb.AgreeRequestsMessage) (*clientPb.AgreementResult, error) {
	// 동의한다는 메시지를 전달
	channelPayments := in.ChannelPayments

	var channelIds []C.uint
	var amount []C.int

	// Extract Data
	for _, channelPayment := range channelPayments.ChannelPayments{
		channelIds = append(channelIds, C.uint(uint32(channelPayment.ChannelId)))
		amount = append(amount, C.int(int32(channelPayment.Amount)))
	}

	//void ecall_go_pre_update_w(unsigned char *msg, unsigned char *signature, unsigned char **original_msg, unsigned char **output);
	convertedOriginalMsg, convertedSignatureMsg := util.ConvertByteToPointer(in.OriginalMessage, in.Signature)

	var originalMsg *C.uchar
	var signature *C.uchar
	C.ecall_go_pre_update_w(convertedOriginalMsg, convertedSignatureMsg, originalMsg, signature)

	originalMessageStr, signatureStr := util.ConvertPointerToByte(originalMsg, signature)

	return &clientPb.AgreementResult{PaymentNumber: in.PaymentNumber, Result: true, OriginalMessage: originalMessageStr, Signature: signatureStr}, nil
}

func (s *ClientGrpc) UpdateRequest(ctx context.Context, in *clientPb.UpdateRequestsMessage) (*clientPb.UpdateResult, error) {
	// 채널 정보를 업데이트 한다던지 잔액을 변경.
	channelPayments := in.ChannelPayments

	var channelIds []C.uint
	var amount []C.int

	// Extract Data
	for _, channelPayment := range channelPayments.ChannelPayments{
		channelIds = append(channelIds, C.uint(uint32(channelPayment.ChannelId)))
		amount = append(amount, C.int(int32(channelPayment.Amount)))
	}

	//void ecall_go_post_update_w(unsigned char *msg, unsigned char *signature, unsigned char **original_msg, unsigned char **output);
	convertedOriginalMsg, convertedSignatureMsg := util.ConvertByteToPointer(in.OriginalMessage, in.Signature)
	var originalMsg *C.uchar
	var signature *C.uchar
	C.ecall_go_pre_update_w(convertedOriginalMsg, convertedSignatureMsg, originalMsg, signature)

	originalMessageStr, signatureStr := util.ConvertPointerToByte(originalMsg, signature)
	C.ecall_go_post_update_w()

	return &clientPb.UpdateResult{PaymentNumber: in.PaymentNumber, Result: true, OriginalMessage: originalMessageStr, Signature: signatureStr}, nil
}

func (s *ClientGrpc) ConfirmPayment(ctx context.Context, in *clientPb.ConfirmRequestsMessage) (*clientPb.Result, error) {
	log.Println("----ConfirmPayment Request Receive----")

	convertedOriginalMsg, convertedSignatureMsg := util.ConvertByteToPointer(in.OriginalMessage, in.Signature)
	C.ecall_go_idle_w(convertedOriginalMsg, convertedSignatureMsg)
	log.Println("----ConfirmPayment Request End----")

	fmt.Println(C.ecall_get_balance_w(C.uint(1)))
	fmt.Println(C.ecall_get_balance_w(C.uint(2)))
	fmt.Println(time.Since(controller.ExecutionTime))

	return &clientPb.Result{PaymentNumber: in.PaymentNumber, Result: true}, nil
}

func (s *ClientGrpc) DirectChannelPayment(ctx context.Context, in *clientPb.ChannelPayment) (*clientPb.DirectPaymentResult, error) {
	log.Println("----Direct Channel Payment Request Receive----")

	originalMessagePointer, signaturePointer := util.ConvertByteToPointer(in.OriginalMessage, in.Signature)

	var replyMessage *C.uchar
	var replySignature *C.uchar

	C.ecall_paid_w(originalMessagePointer, signaturePointer, &replyMessage, &replySignature)
	log.Println("----Direct Channel Payment Request End----")

	convertedReplyMessage, convertedReplySignature := util.ConvertPointerToByte(replyMessage, replySignature)

	return &clientPb.DirectPaymentResult{Result: true, ReplyMessage: convertedReplyMessage, ReplySignature: convertedReplySignature}, nil
}
