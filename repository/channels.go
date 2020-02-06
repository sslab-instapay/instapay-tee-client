package repository

/*
#cgo CPPFLAGS: -I/home/xiaofo/sgxsdk/include -I/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client
#cgo LDFLAGS: -L/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client -ltee

#include "app.h"
*/
import "C"

import (
	"github.com/sslab-instapay/instapay-tee-client/model"
	"fmt"
	"unsafe"
	"reflect"
)

func GetAllChannelsLockedBalance() (int64, error) {

	var ochs unsafe.Pointer
	var lockedBalances int64

	ochs = C.ecall_get_open_channels_w()
	channelSize := 68
	channelSlice := (*[1 << 30]C.channel)(unsafe.Pointer(ochs))[:channelSize:channelSize]

	openChannelNumbers := int(C.ecall_get_num_open_channels_w())

	for i := 0; i < openChannelNumbers; i++ {
		lockedBalances += int64(channelSlice[i].m_locked_balance)
	}

	return lockedBalances, nil
}

func GetClosedChannelList() ([]model.Channel, error) {

	var channelList []model.Channel
	var cchs unsafe.Pointer

	cchs = C.ecall_get_closed_channels_w()
	channelStructSize := 68
	channelSlice := (*[1 << 30]C.channel)(unsafe.Pointer(cchs))[:channelStructSize:channelStructSize]

	closedChannelNumbers := int(C.ecall_get_num_closed_channels_w())

	for i := 0; i < closedChannelNumbers; i++ {
		var channel model.Channel
		channel.ChannelId = int64(channelSlice[i].m_id)
		if channelSlice[i].m_is_in == 0 {
			channel.Type = model.IN
		} else {
			channel.Type = model.OUT
		}
		switch channelSlice[i].m_status {
		case 0:
			channel.Status = model.PENDING
		case 1:
			channel.Status = model.IDLE
		case 2:
			channel.Status = model.PRE_UPDATE
		case 3:
			channel.Status = model.POST_UPDATE
		case 4:
			channel.Status = model.CLOSED
		}
		channel.MyDeposit = int(channelSlice[i].m_my_deposit)
		channel.OtherDeposit = int(channelSlice[i].m_other_deposit)
		channel.MyBalance = int(channelSlice[i].m_balance)
		channel.LockedBalance = int(channelSlice[i].m_locked_balance)

		var sig *C.uchar = &(channelSlice[i].m_my_addr[0])
		hdr := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(sig)),
			Len:  int(20),
			Cap:  int(20),
		}
		s := *(*[]C.uchar)(unsafe.Pointer(&hdr))
		var myAddress string
		myAddress = fmt.Sprintf("%02x", s)
		channel.MyAddress = "0x" + myAddress

		var sig1 *C.uchar = &(channelSlice[i].m_other_addr[0])
		hdr1 := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(sig1)),
			Len:  int(20),
			Cap:  int(20),
		}
		s1 := *(*[]C.uchar)(unsafe.Pointer(&hdr1))
		var otherAddress string
		otherAddress = fmt.Sprintf("%02x", s1)
		channel.OtherAddress = "0x" + otherAddress
		channelList = append(channelList, channel)
	}
	return channelList, nil
}

func GetOpenedChannelList() ([]model.Channel, error) {

	var channelList []model.Channel
	var ochs unsafe.Pointer

	ochs = C.ecall_get_open_channels_w()
	channelSize := 68
	channelSlice := (*[1 << 30]C.channel)(unsafe.Pointer(ochs))[:channelSize:channelSize]

	openChannelNumbers := int(C.ecall_get_num_open_channels_w())

	for i := 0; i < openChannelNumbers; i++ {
		var channel model.Channel
		channel.ChannelId = int64(channelSlice[i].m_id)
		if channelSlice[i].m_is_in == 0 {
			channel.Type = model.IN
		} else {
			channel.Type = model.OUT
		}
		switch channelSlice[i].m_status {
		case 0:
			channel.Status = model.PENDING
		case 1:
			channel.Status = model.IDLE
		case 2:
			channel.Status = model.PRE_UPDATE
		case 3:
			channel.Status = model.POST_UPDATE
		case 4:
			channel.Status = model.CLOSED
		}
		channel.MyDeposit = int(channelSlice[i].m_my_deposit)
		channel.OtherDeposit = int(channelSlice[i].m_other_deposit)
		channel.MyBalance = int(channelSlice[i].m_balance)
		channel.LockedBalance = int(channelSlice[i].m_locked_balance)

		var sig *C.uchar = &(channelSlice[i].m_my_addr[0])
		hdr := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(sig)),
			Len:  int(20),
			Cap:  int(20),
		}
		s := *(*[]C.uchar)(unsafe.Pointer(&hdr))
		var myAddress string
		myAddress = fmt.Sprintf("%02x", s)
		channel.MyAddress = "0x" + myAddress

		var sig1 *C.uchar = &(channelSlice[i].m_other_addr[0])
		hdr1 := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(sig1)),
			Len:  int(20),
			Cap:  int(20),
		}
		s1 := *(*[]C.uchar)(unsafe.Pointer(&hdr1))
		var otherAddress string
		otherAddress = fmt.Sprintf("%02x", s1)
		channel.OtherAddress = "0x" + otherAddress
		channelList = append(channelList, channel)
	}
	return channelList, nil
}

func GetChannelById(channelId int) (model.Channel, error) {

	c1 := C.ecall_get_channel_info_w(C.uint(channelId))

	cvtd := (*C.channel)(unsafe.Pointer(c1))

	var channel model.Channel
	channel.ChannelId = int64(cvtd.m_id)
	if cvtd.m_is_in == 0 {
		channel.Type = model.IN
	} else {
		channel.Type = model.OUT
	}
	switch cvtd.m_status {
	case 0:
		channel.Status = model.PENDING
	case 1:
		channel.Status = model.IDLE
	case 2:
		channel.Status = model.PRE_UPDATE
	case 3:
		channel.Status = model.POST_UPDATE
	case 4:
		channel.Status = model.CLOSED
	}
	channel.MyDeposit = int(cvtd.m_my_deposit)
	channel.OtherDeposit = int(cvtd.m_other_deposit)
	channel.MyBalance = int(cvtd.m_balance)
	channel.LockedBalance = int(cvtd.m_locked_balance)

	var sig *C.uchar = &(cvtd.m_my_addr[0])
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(sig)),
		Len:  int(20),
		Cap:  int(20),
	}
	s := *(*[]C.uchar)(unsafe.Pointer(&hdr))
	var myAddress string
	myAddress = fmt.Sprintf("%02x", s)
	channel.MyAddress = "0x" + myAddress

	var sig1 *C.uchar = &(cvtd.m_other_addr[0])
	hdr1 := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(sig1)),
		Len:  int(20),
		Cap:  int(20),
	}
	s1 := *(*[]C.uchar)(unsafe.Pointer(&hdr1))
	var otherAddress string
	otherAddress = fmt.Sprintf("%02x", s1)
	channel.OtherAddress = "0x" + otherAddress
	return channel, nil
}
