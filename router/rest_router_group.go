package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sslab-instapay/instapay-tee-client/controller"
)

// 라우터 등록 코드
func RegisterRestRouter(router *gin.Engine) {

	accountRouter := router.Group("accounts")
	{
		accountRouter.GET("list", controller.AccountInformationHandler)
		accountRouter.POST("payment", controller.OnchainPaymentHandler)
	}
	walletRouter := router.Group("wallets")
	{
		walletRouter.GET("", controller.GetWalletInformationHandler)
	}
	channelRouter := router.Group("channels")
	{
		channelRouter.GET("list", controller.GetChannelListHandler)
	}
	channelRequestRouter := router.Group("channels/requests")
	{
		channelRequestRouter.POST("open", controller.OpenChannelHandler)
		channelRequestRouter.POST("direct", controller.DirectPayChannelHandler)
		channelRequestRouter.POST("close", controller.CloseChannelHandler)
		channelRequestRouter.POST("server", controller.PaymentToServerChannelHandler)
		channelRequestRouter.POST("eject", controller.EjectChannelHandler)
	}
}
