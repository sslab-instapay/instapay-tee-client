package model

import (
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

type Account struct {
	PublicKeyAddress string    `json:"publicKeyAddress"`
	Balance          big.Float `json:"balance"`
}

type AccountDTO struct {
	PublicKeyAddress string `json:"address"`
	Balance          int  `json:"balance"`
	OffChainDeposit  int  `json:"offchainDeposit"`
	OffChainBalance  int  `json:"offchainBalance"`
}

type ChannelStatus string

const (
	// 0, 1, 2, 3
	PENDING     ChannelStatus = "PENDING"
	IDLE                      = "IDLE"
	PRE_UPDATE                = "PRE_UPDATE"
	POST_UPDATE               = "POST_UPDATE"
	CLOSED                    = "CLOSED"
)

type ChannelType string

const (
	IN  ChannelType = "IN"
	OUT             = "OUT"
)

type PaymentData struct {
	PaymentNumber int64 `bson:"paymentNumber"`
	ChannelId     int64 `bson:"channelId"`
	Amount        int64 `bson:"amount"`
}

type Channel struct {
	ChannelId     int64         `bson:"channelId" json:"channelId"`
	Type          ChannelType   `bson:"channelType" json:"channelType"`
	Status        ChannelStatus `bson:"channelStatus" json:"channelStatus"`
	MyAddress     string        `bson:"myAddress" json:"myAddress"`
	MyDeposit     int         `bson:"myDeposit" json:"myDeposit"`
	OtherDeposit  int         `bson:"otherDeposit" json:"otherDeposit"`
	MyBalance     int         `bson:"myBalance" json:"myBalance"`
	LockedBalance int         `bson:"lockedBalance" json:"lockedBalance"`
	OtherAddress  string        `bson:"otherAddress" json:"otherAddress"`
	OtherIp       string        `bson:"otherIp" json:"otherIp"`
	OtherPort     int           `bson:"otherPort" json:"otherPort"`
}

type CreateChannelEvent struct {
	Id       int64
	Owner    common.Address
	Receiver common.Address
	Deposit  int64
}

type CloseChannelEvent struct {
	Id          int64
	Ownerbal    int64
	Receiverbal int64
}

type EjectEvent struct {
	Pn              int64
	Registeredstage int
}

type PeerInformations struct {
	PeerInformationList []PeerInformation `json:"peer_informations"`
}

type PeerInformation struct {
	PublicKeyAddress string `json:"public_key_address"`
	IpAddress        string `json:"ip_address"`
	GrpcPort         int    `json:"grpc_port"`
}
