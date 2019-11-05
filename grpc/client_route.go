package grpc

/*
#cgo CPPFLAGS: -I/home/xiaofo/sgxsdk/include -I./untrusted -I./include
#cgo LDFLAGS: -L. -ltee
#include "untrusted/app.h"
*/
import "C"

import (
	"context"
	clientPb "github.com/sslab-instapay/instapay-go-client/proto/client"
		"log"
	)

type ClientGrpc struct {
}

func (s *ClientGrpc) AgreementRequest(ctx context.Context, in *clientPb.AgreeRequestsMessage) (*clientPb.Result, error) {
	// 동의한다는 메시지를 전달
	channelPayments := in.ChannelPayments

	var goChannelIds []int64
	var goAmounts []int64

	// Extract Data
	for _, channelPayment := range channelPayments.ChannelPayments{
		goChannelIds = append(goChannelIds, channelPayment.ChannelId)
		goAmounts = append(goAmounts, channelPayment.Amount)
	}

	paymentNum := C.uint(in.PaymentNumber)
	channelIds := []C.uint{goChannelIds}
	amount := []C.int{goAmounts}
	size := C.uint(len(goChannelIds))
	C.ecall_go_pre_update_w(paymentNum, &channelIds[0], &amount[0], size)

	return &clientPb.Result{PaymentNumber: in.PaymentNumber, Result: true}, nil
}

func (s *ClientGrpc) UpdateRequest(ctx context.Context, in *clientPb.UpdateRequestsMessage) (*clientPb.Result, error) {
	// 채널 정보를 업데이트 한다던지 잔액을 변경.
	channelPayments := in.ChannelPayments

	var goChannelIds []int64
	var goAmounts []int64

	// Extract Data
	for _, channelPayment := range channelPayments.ChannelPayments{
		goChannelIds = append(goChannelIds, channelPayment.ChannelId)
		goAmounts = append(goAmounts, channelPayment.Amount)
	}

	paymentNum := C.uint(in.PaymentNumber)
	channelIds := []C.uint{goChannelIds}
	amount := []C.int{goAmounts}
	size := C.uint(len(goChannelIds))

	C.ecall_go_post_update_w(paymentNum, &channelIds[0], &amount[0], size)
	return &clientPb.Result{PaymentNumber: in.PaymentNumber, Result: true}, nil
}

func (s *ClientGrpc) ConfirmPayment(ctx context.Context, in *clientPb.ConfirmRequestsMessage) (*clientPb.Result, error) {
	log.Println("----ConfirmPayment Request Receive----")
	C.ecall_go_idle_w(in.PaymentNumber)
	log.Println("----ConfirmPayment Request End----")

	return &clientPb.Result{PaymentNumber: in.PaymentNumber, Result: true}, nil
}
