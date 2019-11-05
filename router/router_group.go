package router

import (
	"github.com/gin-gonic/gin"
		"github.com/sslab-instapay/instapay-go-client/controller"
)

func RegisterChannelRouter(router *gin.Engine){

	channelRouter := router.Group("channels/requests")
	{
		channelRouter.POST("open", controller.OpenChannelHandler)

		//TODO 데모 이후 추가.
		//channelRouter.POST("deposit", func(context *gin.Context) {
		//		//	context.JSON(http.StatusOK, controller.DepositChannelHandler)
		//		//})

		channelRouter.POST("direct", controller.DirectPayChannelHandler)

		channelRouter.POST("close", controller.CloseChannelHandler)

		channelRouter.POST("server", controller.PaymentToServerChannelHandler)
	}
}





