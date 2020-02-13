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
	// "github.com/sslab-instapay/instapay-tee-client/controller"
	"log"
	// "fmt"
	// "time"
	"unsafe"
	"reflect"
)

type ClientGrpc struct {
}

func (s *ClientGrpc) AgreementRequest(ctx context.Context, in *clientPb.AgreeRequestsMessage) (*clientPb.AgreementResult, error) {

	convertedOriginalMsg, convertedSignatureMsg := convertByteToPointer(in.OriginalMessage, in.Signature)

	var originalMsg *C.uchar
	var signature *C.uchar
	C.ecall_go_pre_update_w(convertedOriginalMsg, convertedSignatureMsg, &originalMsg, &signature)

	originalMessageStr, signatureStr := convertPointerToByte(originalMsg, signature)

	defer C.free(unsafe.Pointer(originalMsg))
	defer C.free(unsafe.Pointer(signature))

	return &clientPb.AgreementResult{Result: true, OriginalMessage: originalMessageStr, Signature: signatureStr}, nil
}

func (s *ClientGrpc) UpdateRequest(ctx context.Context, in *clientPb.UpdateRequestsMessage) (*clientPb.UpdateResult, error) {
	// 채널 정보를 업데이트 한다던지 잔액을 변경.
	convertedOriginalMsg, convertedSignatureMsg := convertByteToPointer(in.OriginalMessage, in.Signature)

	var originalMsg *C.uchar
	var signature *C.uchar
	C.ecall_go_post_update_w(convertedOriginalMsg, convertedSignatureMsg, &originalMsg, &signature)

	originalMessageStr, signatureStr := convertPointerToByte(originalMsg, signature)
	defer C.free(unsafe.Pointer(originalMsg))
	defer C.free(unsafe.Pointer(signature))

	return &clientPb.UpdateResult{Result: true, OriginalMessage: originalMessageStr, Signature: signatureStr}, nil
}

func (s *ClientGrpc) ConfirmPayment(ctx context.Context, in *clientPb.ConfirmRequestsMessage) (*clientPb.ConfirmResult, error) {
	log.Println("----ConfirmPayment Request Receive----")

	convertedOriginalMsg, convertedSignatureMsg := convertByteToPointer(in.OriginalMessage, in.Signature)
	C.ecall_go_idle_w(convertedOriginalMsg, convertedSignatureMsg)
	log.Println("----ConfirmPayment Request End----")

	// fmt.Println(C.ecall_get_balance_w(C.uint(1)))
	// fmt.Println(C.ecall_get_balance_w(C.uint(2)))
	// fmt.Println(time.Since(controller.ExecutionTime))

	return &clientPb.ConfirmResult{Result: true}, nil
}

func (s *ClientGrpc) DirectChannelPayment(ctx context.Context, in *clientPb.DirectChannelPaymentMessage) (*clientPb.DirectPaymentResult, error) {
	log.Println("----Direct Channel Payment Request Receive----")

	log.Println("--- Start Byte to Pointer ---")
	originalMessagePointer, signaturePointer := convertByteToPointer(in.OriginalMessage, in.Signature)
	log.Println("--- End Byte to Pointer ---")
	var replyMessage *C.uchar
	var replySignature *C.uchar

	C.ecall_paid_w(originalMessagePointer, signaturePointer, &replyMessage, &replySignature)
	log.Println("----Direct Channel Payment Request End----")

	convertedReplyMessage, convertedReplySignature := convertPointerToByte(replyMessage, replySignature)

	defer C.free(unsafe.Pointer(replyMessage))
	defer C.free(unsafe.Pointer(replySignature))

	return &clientPb.DirectPaymentResult{Result: true, ReplyMessage: convertedReplyMessage, ReplySignature: convertedReplySignature}, nil
}

func convertByteToPointer(originalMsg []byte, signature []byte) (*C.uchar, *C.uchar){

	var uOriginal [44]C.uchar
	var uSignature [65]C.uchar

	for i := 0; i < 44; i++{
		uOriginal[i] = C.uchar(originalMsg[i])
	}

	for i := 0; i < 65; i++{
		uSignature[i] = C.uchar(signature[i])
	}

	cOriginalMsg := (*C.uchar)(unsafe.Pointer(&uOriginal[0]))
	cSignature := (*C.uchar)(unsafe.Pointer(&uSignature[0]))

	return cOriginalMsg, cSignature
}

func convertPointerToByte(originalMsg *C.uchar, signature *C.uchar)([]byte, []byte){

	var returnMsg []byte
	var returnSignature []byte

	replyMsgHdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(originalMsg)),
		Len: int(44),
		Cap: int(44),
	}
	replyMsgS := *(*[]C.uchar)(unsafe.Pointer(&replyMsgHdr))

	replySigHdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(signature)),
		Len: int(65),
		Cap: int(65),
	}
	replySigS := *(*[]C.uchar)(unsafe.Pointer(&replySigHdr))

	for i := 0; i < 44; i++{
		returnMsg = append(returnMsg, byte(replyMsgS[i]))
	}

	for i := 0; i < 65; i++{
		returnSignature = append(returnSignature, byte(replySigS[i]))
	}

	return returnMsg, returnSignature
}