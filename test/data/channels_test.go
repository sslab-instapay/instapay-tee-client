package data_test

import (
	"fmt"
	"testing"
	"log"
	"github.com/sslab-instapay/instapay-go-client/repository"
	"github.com/sslab-instapay/instapay-go-client/model"
	"github.com/sslab-instapay/instapay-go-client/config"
	"os"
)

func TestGetChannelList(t *testing.T){
	channelList, err := repository.GetChannelList()

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(channelList)
}

func TestGetClosedChannelList(t *testing.T){
	channelList, err := repository.GetClosedChannelList()

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(channelList)
}

func TestGetOpenedChannelList(t *testing.T){
	channelList, err := repository.GetOpenedChannelList()

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(channelList)
}

func TestGetChannelById(t *testing.T){
	os.Setenv("database_name", "instapay-client")
	channel, err := repository.GetChannelById(2)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(channel)
}

func TestUpdateChannel(t *testing.T){
	os.Setenv("database_name", "instapay-client")
	channel, err := repository.GetChannelById(5)
	if err != nil{
		log.Fatal(err)
	}
	
	channel.Status = model.CLOSED
	updatedChannel, err := repository.UpdateChannel(channel)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updatedChannel)
}

func TestInsertChannel(t *testing.T){
	channel := model.Channel{ChannelId: 1, Type: model.OUT,
		Status: model.IDLE, MyAddress: config.GetAccountConfig().PublicKeyAddress,
		MyBalance: 100, MyDeposit: 100, OtherDeposit: 0, OtherAddress: "0xD03A2CC08755eC7D75887f0997195654b928893e"}

	insertedChannel, err := repository.InsertChannel(channel)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(insertedChannel)
}

func TestGetAllChannelsLockedBalance(t *testing.T){

	lockedBalance, err := repository.GetAllChannelsLockedBalance()

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(lockedBalance)
}

func TestInsertPaymentData(t *testing.T){
	os.Setenv("database_name", "instapay-client")
	repository.InsertPaymentData(model.PaymentData{PaymentNumber: 3, ChannelId: 3, Amount: 9})

	fmt.Println(repository.GetPaymentDatasByPaymentNumber(3))
}

func TestGetPaymentDataByPaymentId(t *testing.T){

	os.Setenv("database_name", "instapay-client")
	paymentData, err := repository.GetPaymentDatasByPaymentNumber(1)

	if err != nil{
		log.Println("HO")
		log.Fatal(err)
	}

	fmt.Println(paymentData)
}

func TestFindPaymentData(t *testing.T){

	os.Setenv("database_name", "instapay-client")
	result, err := repository.FindPaymentData(model.PaymentData{PaymentNumber:1, ChannelId:2, Amount:5})

	if err != nil{
		log.Println("HO")
		log.Fatal(err)
	}

	fmt.Println(result)
}



