package repository

import "C"
import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"github.com/sslab-instapay/instapay-tee-client/model"
	"github.com/sslab-instapay/instapay-tee-client/db"
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

	openChannelNumbers := C.ecall_get_num_open_channels()

	for i := 0; i < openChannelNumbers; i++{
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

	closedChannelNumbers := C.ecall_get_num_closed_channels_w()

	for i := 0; i < closedChannelNumbers; i++{
		var channel model.Channel
		channel.ChannelId = int64(channelSlice[i].m_id)
		if channelSlice[i].m_is_in == 0{
			channel.Type = model.IN
		}else{
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
		channel.MyDeposit = channelSlice[i].m_my_deposit
		channel.OtherDeposit = channelSlice[i].m_other_deposit
		channel.MyBalance = channelSlice[i].m_my_balance
		channel.LockedBalance = channelSlice[i].m_locked_balance

		var sig *C.uchar = &(channelSlice[i].m_my_addr[0])
		hdr := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(sig)),
			Len: int(20),
			Cap: int(20),
		}
		s := *(*[]C.uchar)(unsafe.Pointer(&hdr))
		var myAddress string
		myAddress = fmt.Sprintf("%02x", s)
		channel.MyAddress = "0x" + myAddress

		var sig1 *C.uchar = &(channelSlice[i].m_other_addr[0])
		hdr1 := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(sig1)),
			Len: int(20),
			Cap: int(20),
		}
		s1 := *(*[]C.uchar)(unsafe.Pointer(&hdr1))
		var otherAddress string
		otherAddress = fmt.Sprint("%02x", s1)
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

	openChannelNumbers := C.ecall_get_num_open_channels()

	for i := 0; i < openChannelNumbers; i++{
		var channel model.Channel
		channel.ChannelId = int64(channelSlice[i].m_id)
		if channelSlice[i].m_is_in == 0{
			channel.Type = model.IN
		}else{
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
		channel.MyDeposit = channelSlice[i].m_my_deposit
		channel.OtherDeposit = channelSlice[i].m_other_deposit
		channel.MyBalance = channelSlice[i].m_my_balance
		channel.LockedBalance = channelSlice[i].m_locked_balance

		var sig *C.uchar = &(channelSlice[i].m_my_addr[0])
		hdr := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(sig)),
			Len: int(20),
			Cap: int(20),
		}
		s := *(*[]C.uchar)(unsafe.Pointer(&hdr))
		var myAddress string
		myAddress = fmt.Sprintf("%02x", s)
		channel.MyAddress = "0x" + myAddress

		var sig1 *C.uchar = &(channelSlice[i].m_other_addr[0])
		hdr1 := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(sig1)),
			Len: int(20),
			Cap: int(20),
		}
		s1 := *(*[]C.uchar)(unsafe.Pointer(&hdr1))
		var otherAddress string
		otherAddress = fmt.Sprint("%02x", s1)
		channel.OtherAddress = "0x" + otherAddress
		channelList = append(channelList, channel)
	}
	return channelList, nil
}

func GetChannelById(channelId int64) (model.Channel, error) {

	database, err := db.GetDatabase()
	if err != nil {
		return model.Channel{}, err
	}

	filter := bson.M{
		"channelId": channelId,
	}

	collection := database.Collection("channels")

	channel := model.Channel{}
	singleRecord := collection.FindOne(context.TODO(), filter)
	if err := singleRecord.Decode(&channel); err != nil {
		log.Println(err)
	}
	return channel, nil
}
