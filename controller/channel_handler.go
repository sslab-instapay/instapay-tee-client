package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sslab-instapay/instapay-tee-client/repository"
	"log"
	"strconv"
	"github.com/sslab-instapay/instapay-tee-client/service"
	"github.com/sslab-instapay/instapay-tee-client/config"
	serverPb "github.com/sslab-instapay/instapay-tee-client/proto/server"
	"google.golang.org/grpc"
	"time"
	"context"
	"github.com/sslab-instapay/instapay-tee-client/model"
)

var ExecutionTime time.Time

func OpenChannelHandler(ctx *gin.Context) {

	otherAddress := ctx.PostForm("other_addr")
	deposit, _ := strconv.Atoi(ctx.PostForm("deposit"))

	txHash, err := service.SendOpenChannelTransaction(deposit, otherAddress)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"result": "success", "txHash": txHash})
	}
}

func CloseChannelHandler(ctx *gin.Context) {
	channelIdParam := ctx.PostForm("channelId")
	log.Println(channelIdParam)
	channelId, _ := strconv.Atoi(channelIdParam)
	log.Println(channelId)

	service.SendCloseChannelTransaction(int64(channelId))

	ctx.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func EjectChannelHandler(ctx *gin.Context){

}

// TODO 데모 시나리오 이후 구현
func DepositChannelHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func DirectPayChannelHandler(context *gin.Context) {
	//channelId := context.PostForm("ch_id")
	//amount := context.PostForm("amount")

	//conn, err := grpc.Dial(config.EthereumConfig["serverAddr"], grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	//c := pb.NewGreeterClient(conn)
	//
	//// Contact the server and print out its response.
	//name := defaultName
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.GetMessage())

	context.JSON(http.StatusOK, gin.H{"message": "Channel"})
}

func PaymentToServerChannelHandler(ctx *gin.Context) {

	otherAddress := ctx.PostForm("addr")
	amount, err := strconv.Atoi(ctx.PostForm("amount"))
	if err != nil {
		log.Println(err)
	}

	myAddress := config.GetAccountConfig().PublicKeyAddress
	connection, err := grpc.Dial(config.EthereumConfig["serverGrpcHost"]+":"+config.EthereumConfig["serverGrpcPort"], grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	defer connection.Close()
	client := serverPb.NewServerClient(connection)

	clientContext, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ExecutionTime = time.Now()
	r, err := client.PaymentRequest(clientContext, &serverPb.PaymentRequestMessage{From: myAddress, To: otherAddress, Amount: int64(amount)})
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	log.Println(r.GetResult())

	ctx.JSON(http.StatusOK, gin.H{"sendAddress": otherAddress, "amount": amount})
}

func GetChannelListHandler(ctx *gin.Context) {

	channelList, err := repository.GetChannelList()
	if err != nil {
		log.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"channels": channelList,
	})
}

func GetWalletInformationHandler(ctx *gin.Context) {

	account := config.GetAccountConfig()

	tempBalance, err := service.GetBalance()
	var balance int64
	if err != nil {
		log.Println(err)
	}

	balance, _ = tempBalance.Int64()

	offchainDeposit, err := repository.GetAllDepositValue()
	if err != nil {
		log.Println(err)
	}

	offchainBalance, err := repository.GetOffChainBalance()
	if err != nil {
		log.Println(err)
	}

	accountDto := model.AccountDTO{
		PublicKeyAddress: account.PublicKeyAddress,
		Balance:          balance,
		OffChainDeposit:  offchainDeposit,
		OffChainBalance:  offchainBalance,
	}

	openedChannelList, err := repository.GetOpenedChannelList()
	if err != nil {
		log.Println(err)
	}

	inChannelList := make([]model.Channel, 0)
	outChannelList := make([]model.Channel, 0)

	for _, channel := range openedChannelList {
		if channel.Type == model.IN {
			inChannelList = append(inChannelList, channel)
		} else {
			outChannelList = append(outChannelList, channel)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"account": accountDto, "inChannelList": inChannelList, "outChannelList": outChannelList,
	})

}
