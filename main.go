package main

/*
#cgo CPPFLAGS: -I/home/xiaofo/sgxsdk/include -I/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client
#cgo LDFLAGS: -L/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client -ltee

#include "app.h"
*/
import "C"

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sslab-instapay/instapay-tee-client/config"
	instapayGrpc "github.com/sslab-instapay/instapay-tee-client/grpc"
	clientPb "github.com/sslab-instapay/instapay-tee-client/proto/client"
	"github.com/sslab-instapay/instapay-tee-client/router"
	"github.com/sslab-instapay/instapay-tee-client/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
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
	// TODO seal 코드 들어가기
	config.SetAccountConfig("0x222")
}
