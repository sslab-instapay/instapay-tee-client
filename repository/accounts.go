package repository

import "C"
import (
							"unsafe"
)

func GetAllDepositValue() (int64, error) {

	var ochs unsafe.Pointer
	var depositValue int64

	ochs = C.ecall_get_open_channels_w()
	channelSize := 68
	channelSlice := (*[1 << 30]C.channel)(unsafe.Pointer(ochs))[:channelSize:channelSize]

	openChannelNumbers := C.ecall_get_num_open_channels()

	for i := 0; i < openChannelNumbers; i++{
		depositValue += int64(channelSlice[i].m_my_deposit)
	}

	return depositValue, nil
}

func GetOffChainBalance() (int64, error) {

	var ochs unsafe.Pointer
	var offchainBalance int64

	ochs = C.ecall_get_open_channels_w()
	channelSize := 68
	channelSlice := (*[1 << 30]C.channel)(unsafe.Pointer(ochs))[:channelSize:channelSize]

	openChannelNumbers := C.ecall_get_num_open_channels()

	for i := 0; i < openChannelNumbers; i++{
		offchainBalance += int64(channelSlice[i].m_my_balance)
	}
	return offchainBalance, nil
}
