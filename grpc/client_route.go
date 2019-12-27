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
)

type ClientGrpc struct {
}

func (s *ClientGrpc) AgreementRequest(ctx context.Context, in *clientPb.AgreeRequestsMessage) (*clientPb.Result, error) {
	// 동의한다는 메시지를 전달
	channelPayments := in.ChannelPayments

	var channelIds []C.uint
	var amount []C.int

	// Extract Data
	for _, channelPayment := range channelPayments.ChannelPayments{
		channelIds = append(channelIds, C.uint(uint32(channelPayment.ChannelId)))
		amount = append(amount, C.int(int32(channelPayment.Amount)))
	}

	paymentNum := C.uint(uint32(in.PaymentNumber))
	size := C.uint(len(channelIds))
	C.ecall_go_pre_update_w(paymentNum, &channelIds[0], &amount[0], size)

	return &clientPb.Result{PaymentNumber: in.PaymentNumber, Result: true}, nil
}

func (s *ClientGrpc) UpdateRequest(ctx context.Context, in *clientPb.UpdateRequestsMessage) (*clientPb.Result, error) {
	// 채널 정보를 업데이트 한다던지 잔액을 변경.
	channelPayments := in.ChannelPayments

	var channelIds []C.uint
	var amount []C.int

	// Extract Data
	for _, channelPayment := range channelPayments.ChannelPayments{
		channelIds = append(channelIds, C.uint(uint32(channelPayment.ChannelId)))
		amount = append(amount, C.int(int32(channelPayment.Amount)))
	}

	paymentNum := C.uint(uint32(in.PaymentNumber))
	size := C.uint(uint32(len(channelIds)))
	C.ecall_go_post_update_w(paymentNum, &channelIds[0], &amount[0], size)

	return &clientPb.Result{PaymentNumber: in.PaymentNumber, Result: true}, nil
}

func (s *ClientGrpc) ConfirmPayment(ctx context.Context, in *clientPb.ConfirmRequestsMessage) (*clientPb.Result, error) {
	log.Println("----ConfirmPayment Request Receive----")
	C.ecall_go_idle_w(C.uint(uint32(in.PaymentNumber)))
	log.Println("----ConfirmPayment Request End----")

	fmt.Println(C.ecall_get_balance_w(C.uint(1)))
	fmt.Println(C.ecall_get_balance_w(C.uint(2)))
	fmt.Println(time.Since(controller.ExecutionTime))

	return &clientPb.Result{PaymentNumber: in.PaymentNumber, Result: true}, nil
}
