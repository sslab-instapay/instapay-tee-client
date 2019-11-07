package model

import (
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

type Account struct {
	PublicKeyAddress string `json:"publicKeyAddress"`
	PrivateKey       string
	Balance          big.Float `json:"balance"`
}

type AccountDTO struct {
	PublicKeyAddress string `json:"address"`
	Balance          int64  `json:"balance"`
	OffChainDeposit  int64  `json:"offchainDeposit"`
	OffChainBalance  int64  `json:"offchainBalance"`
}

type ChannelStatus string

const (
	// 0, 1, 2, 3
	IDLE        ChannelStatus = "IDLE"
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
	MyDeposit     int64         `bson:"myDeposit" json:"myDeposit"`
	OtherDeposit  int64         `bson:"otherDeposit" json:"otherDeposit"`
	MyBalance     int64         `bson:"myBalance" json:"myBalance"`
	LockedBalance int64         `bson:"lockedBalance" json:"lockedBalance"`
	OtherAddress  string        `bson:"otherAddress" json:"otherAddress"`
	OtherIp       string        `bson:"otherIp" json:"otherIp"`
	OtherPort     int           `bson:"otherPort" json:"otherPort"`
}

type CreateChannelEvent struct {
	Id       int64
	Owner    common.Address
	Receiver common.Address
	Deposit  *big.Int
}

type CloseChannelEvent struct {
	Id          int64
	Ownerbal    *big.Int
	Receiverbal *big.Int
}

type EjectEvent struct {
	Pn              int64
	Registeredstage int
}
