package main

/*
#cgo CPPFLAGS: -I/home/xiaofo/sgxsdk/include -I/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client
#cgo LDFLAGS: -L/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client -ltee

#include "app.h"
*/
import "C"

import (
	instapayGrpc "github.com/sslab-instapay/instapay-tee-client/grpc"
	clientPb "github.com/sslab-instapay/instapay-tee-client/proto/client"
	"net"
	"log"
	"fmt"
	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"
	"github.com/sslab-instapay/instapay-tee-client/router"
	"os"
	"strconv"
		"flag"
	"github.com/sslab-instapay/instapay-tee-client/service"
	"github.com/sslab-instapay/instapay-tee-client/config"
	"github.com/sslab-instapay/instapay-tee-client/repository"
)

func startGrpcServer(){
	log.Println("---Start Grpc Server---")
	grpcPort, err := strconv.Atoi(os.Getenv("grpc_port"))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	clientPb.RegisterClientServer(grpcServer, &instapayGrpc.ClientGrpc{})
	grpcServer.Serve(lis)
}

func startClientWebServer(){

	defaultRouter := gin.Default()
	defaultRouter.LoadHTMLGlob("templates/*")

	defaultRouter.Use(CORSMiddleware())
	router.RegisterRestRouter(defaultRouter)
	router.RegisterChannelRouter(defaultRouter)
	router.RegisterViewRouter(defaultRouter)

	defaultRouter.Run(":" + os.Getenv("port"))
}

func main() {

	C.initialize_enclave()
	LoadDataToTEE()
	portNum := flag.String("port", "3001", "port number")
	grpcPortNum := flag.String("grpc_port", "50001", "grpc_port number")
	databaseName := flag.String("database_name", "instapay-client", "database Name")

	flag.Parse()

	os.Setenv("port", *portNum)
	os.Setenv("grpc_port", *grpcPortNum)
	os.Setenv("database_name", *databaseName)

	go service.ListenContractEvent()
	go startGrpcServer()
	startClientWebServer()

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func LoadDataToTEE(){
	// TODO 데이터 베이스 데이터를 TEE에 로드하자 (Account, Channel)
	account := config.GetAccountConfig()

	pubKey := account.PublicKeyAddress[2:]
	privKey := account.PrivateKey

	teePublicKey := []C.uchar(pubKey)
	teePrivateKey := []C.uchar(privKey)

	C.ecall_preset_account_w(&teePublicKey[0], &teePrivateKey[0])

	channelList, err := repository.GetOpenedChannelList()
	if err != nil{
		log.Println()
	}

	for _, channel := range channelList  {
		myAddress := []C.uchar(channel.MyAddress[2:])
		otherAddress := []C.uchar(channel.OtherAddress[2:])
		otherIpAddress := []C.uchar(channel.OtherIp)
		C.ecall_load_channel_data_w(uint32(channel.ChannelId), channel.Type, channel.Status, &myAddress[0], uint32(channel.MyDeposit), uint32(channel.OtherDeposit), uint32(channel.MyBalance), uint32(channel.LockedBalance), &otherAddress[0], &otherIpAddress[0], uint32(channel.OtherPort))
	}

	log.Println("--- TEE Data Load Successfully!!--- ")

}