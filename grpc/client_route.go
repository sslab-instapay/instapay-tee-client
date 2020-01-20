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
	"unsafe"
	"reflect"
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
	convertedOriginalMsg := &([]C.uchar(in.OriginalMessage)[0])
	convertedSignatureMsg := &([]C.uchar(in.Signature)[0])

	var originalMsg *C.uchar
	var signature *C.uchar
	C.ecall_go_pre_update_w(convertedOriginalMsg, convertedSignatureMsg, originalMsg, signature)

	hdr1 := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(originalMsg)),
		Len:  int(44),
		Cap:  int(44),
	}
	s1 := *(*[]C.uchar)(unsafe.Pointer(&hdr1))

	hdr2 := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(signature)),
		Len:  int(65),
		Cap:  int(65),
	}
	s2 := *(*[]C.uchar)(unsafe.Pointer(&hdr2))
	originalMessageStr := fmt.Sprintf("02x", s1)
	signatureStr := fmt.Sprintf("02x", s2)

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

	// TODO 서명, 메시지 넣은 후 전송
	//void ecall_go_post_update_w(unsigned char *msg, unsigned char *signature, unsigned char **original_msg, unsigned char **output);

	convertedOriginalMsg := &([]C.uchar(in.OriginalMessage)[0])
	convertedSignatureMsg := &([]C.uchar(in.Signature)[0])

	var originalMsg *C.uchar
	var signature *C.uchar
	C.ecall_go_pre_update_w(convertedOriginalMsg, convertedSignatureMsg, originalMsg, signature)

	hdr1 := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(originalMsg)),
		Len:  int(44),
		Cap:  int(44),
	}
	s1 := *(*[]C.uchar)(unsafe.Pointer(&hdr1))

	hdr2 := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(signature)),
		Len:  int(65),
		Cap:  int(65),
	}
	s2 := *(*[]C.uchar)(unsafe.Pointer(&hdr2))
	originalMessageStr := fmt.Sprintf("02x", s1)
	signatureStr := fmt.Sprintf("02x", s2)
	C.ecall_go_post_update_w()

	return &clientPb.UpdateResult{PaymentNumber: in.PaymentNumber, Result: true, OriginalMessage: originalMessageStr, Signature: signatureStr}, nil
}

func (s *ClientGrpc) ConfirmPayment(ctx context.Context, in *clientPb.ConfirmRequestsMessage) (*clientPb.Result, error) {
	log.Println("----ConfirmPayment Request Receive----")
	//void ecall_go_idle_w(unsigned char *msg, unsigned char *signature);
	convertedOriginalMsg := &([]C.uchar(in.OriginalMessage)[0])
	convertedSignatureMsg := &([]C.uchar(in.Signature)[0])
	C.ecall_go_idle_w(convertedOriginalMsg, convertedSignatureMsg)
	log.Println("----ConfirmPayment Request End----")

	fmt.Println(C.ecall_get_balance_w(C.uint(1)))
	fmt.Println(C.ecall_get_balance_w(C.uint(2)))
	fmt.Println(time.Since(controller.ExecutionTime))

	return &clientPb.Result{PaymentNumber: in.PaymentNumber, Result: true}, nil
}

func (s *ClientGrpc) DirectChannelPayment(ctx context.Context, in *clientPb.ChannelPayment) (*clientPb.DirectPaymentResult, error) {
	log.Println("----Direct Channel Payment Request Receive----")
	originalMessage := C.CString(in.OriginalMessage)
	signature := C.CString(in.OriginalMessage)
	defer C.free(unsafe.Pointer(originalMessage))
	defer C.free(unsafe.Pointer(signature))

	var replyMessage *C.uchar
	var replySignature *C.uchar

	C.ecall_paid_w(originalMessage, signature, &replyMessage, &replySignature)
	log.Println("----Direct Channel Payment Request End----")

	fmt.Println(C.ecall_get_balance_w(C.uint(1)))
	fmt.Println(C.ecall_get_balance_w(C.uint(2)))
	fmt.Println(time.Since(controller.ExecutionTime))

	convertedReplyMessage := C.GoString(replyMessage)
	convertedReplySignature := C.GoString(replySignature)

	return &clientPb.DirectPaymentResult{Result: true, ReplyMessage: convertedReplyMessage, ReplySignature: convertedReplySignature}, nil
}
