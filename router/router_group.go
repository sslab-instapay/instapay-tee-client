package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sslab-instapay/instapay-tee-client/controller"
)

func RegisterChannelRouter(router *gin.Engine) {

	channelRouter := router.Group("channels/requests")
	{
		channelRouter.POST("open", controller.OpenChannelHandler)

		channelRouter.POST("direct", controller.DirectPayChannelHandler)

		channelRouter.POST("close", controller.CloseChannelHandler)

		channelRouter.POST("server", controller.PaymentToServerChannelHandler)
	}
}
