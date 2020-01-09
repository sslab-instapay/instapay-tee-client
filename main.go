package main

/*
#cgo CPPFLAGS: -I/home/xiaofo/sgxsdk/include -I/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client
#cgo LDFLAGS: -L/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client -ltee

#include "app.h"
*/
import "C"

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sslab-instapay/instapay-tee-client/config"
	instapayGrpc "github.com/sslab-instapay/instapay-tee-client/grpc"
	clientPb "github.com/sslab-instapay/instapay-tee-client/proto/client"
	"github.com/sslab-instapay/instapay-tee-client/router"
		"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
	"unsafe"
	"flag"
	"github.com/sslab-instapay/instapay-tee-client/service"
)

func main() {
	C.initialize_enclave()
	portNum := flag.String("port", "3001", "port number")
	grpcPortNum := flag.String("grpc_port", "50001", "grpc_port number")
	databaseName := flag.String("database_name", "instapay-client", "database Name")

	flag.Parse()

	os.Setenv("port", *portNum)
	os.Setenv("grpc_port", *grpcPortNum)
	os.Setenv("database_name", *databaseName)
	LoadDataToTEE()
	go service.ListenContractEvent()
	go startGrpcServer()
	startClientWebServer()

}

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
	router.RegisterViewRouter(defaultRouter)

	defaultRouter.Run(":" + os.Getenv("port"))
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
	C.ecall_load_account_data_w()
	C.ecall_load_channel_data_w()

	var paddrs unsafe.Pointer

	paddrs = C.ecall.get_public_addrs_w()
	paddrSize := 20
	paddrSlice := (*[1 << 30]C.address)(unsafe.Pointer(paddrs))[:paddrSize:paddrSize]

	var convertedAddress string
	convertedAddress = fmt.Sprintf("%02x", paddrSlice[0].addr)
	convertedAddress = "0x" + convertedAddress
	fmt.Println("---- Public Key Address ---")
	fmt.Println(convertedAddress)
	config.SetAccountConfig(convertedAddress)
}
