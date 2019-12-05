package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sslab-instapay/instapay-tee-client/controller"
)

// 라우터 등록 코드
func RegisterRestRouter(router *gin.Engine) {

	accountRouter := router.Group("account")
	{
		accountRouter.GET("list", controller.AccountInformationHandler)
		accountRouter.GET("new", controller.CreateAccountHandler)
	}
	walletRouter := router.Group("wallets")
	{
		walletRouter.GET("", controller.GetWalletInformationHandler)
	}
	channelRouter := router.Group("channel")
	{
		channelRouter.GET("list", controller.GetChannelListHandler)
	}
}
