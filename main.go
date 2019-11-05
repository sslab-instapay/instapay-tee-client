package main

import (
	instapayGrpc "github.com/sslab-instapay/instapay-go-client/grpc"
	clientPb "github.com/sslab-instapay/instapay-go-client/proto/client"
	"net"
	"log"
	"fmt"
	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"
	"github.com/sslab-instapay/instapay-go-client/router"
	"os"
	"strconv"
		"flag"
	"github.com/sslab-instapay/instapay-go-client/service"
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
	// os[1] os[2] 로 전역변수 셋팅.

	portNum := flag.String("port", "3001", "port number")
	grpcPortNum := flag.String("grpc_port", "50001", "grpc_port number")
	databaseName := flag.String("database_name", "instapay-client", "database Name")

	flag.Parse()

	os.Setenv("port", *portNum)
	os.Setenv("grpc_port", *grpcPortNum)
	os.Setenv("database_name", *databaseName)
	//
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