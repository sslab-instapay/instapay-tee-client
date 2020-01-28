package repository

import "C"
import (
							"unsafe"
)

func GetAllDepositValue() (int, error) {

	var ochs unsafe.Pointer
	var depositValue int

	ochs = C.ecall_get_open_channels_w()
	channelSize := 68
	channelSlice := (*[1 << 30]C.channel)(unsafe.Pointer(ochs))[:channelSize:channelSize]

	openChannelNumbers := C.ecall_get_num_open_channels_w()

	for i := 0; i < openChannelNumbers; i++{
		depositValue += int(channelSlice[i].m_my_deposit)
	}

	return depositValue, nil
}

func GetOffChainBalance() (int, error) {

	var ochs unsafe.Pointer
	var offchainBalance int

	ochs = C.ecall_get_open_channels_w()
	channelSize := 68
	channelSlice := (*[1 << 30]C.channel)(unsafe.Pointer(ochs))[:channelSize:channelSize]

	openChannelNumbers := int(C.ecall_get_num_open_channels_w())

	for i := 0; i < openChannelNumbers; i++{
		offchainBalance += int(channelSlice[i].m_balance)
	}
	return offchainBalance, nil
}
