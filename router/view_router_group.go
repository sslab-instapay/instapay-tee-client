package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
		"github.com/sslab-instapay/instapay-go-client/repository"
		"github.com/sslab-instapay/instapay-go-client/config"
		)

func RegisterViewRouter(router *gin.Engine) {

	viewRouter := router.Group("templates")
	{
		// account 리스트
		viewRouter.GET("accounts/list", func(context *gin.Context) {
			//gin.H 부분에서 변수 다루는 것.
			account := config.GetAccountConfig()
			context.HTML(http.StatusOK, "account.tmpl", gin.H{
				"account": account,
			})
		})

		// channel 리스트
		viewRouter.GET("channels/list", func(context *gin.Context) {
			channelList, err := repository.GetChannelList()
			if err != nil {
				log.Fatal(err)
			}
			context.HTML(http.StatusOK, "channels.tmpl", gin.H{"channelList": channelList})
		})

		// channel 오픈
		viewRouter.GET("channels/open", func(context *gin.Context) {
			account := config.GetAccountConfig()
			context.HTML(http.StatusOK, "openChannel.tmpl", gin.H{"account": account})
		})

		// 닫힌 channel들
		viewRouter.GET("channels/closed", func(context *gin.Context) {
			closedChannelList, err := repository.GetClosedChannelList()
			if err != nil {
				log.Fatal(err)
			}
			context.HTML(http.StatusOK, "closedChannel.tmpl", gin.H{"closedChannelList": closedChannelList})
		})

		// Pay 페이지
		viewRouter.GET("channels/pay", func(context *gin.Context) {
			channelIdList, err := repository.GetChannelIdList()
			if err != nil {
				log.Fatal(err)
			}

			context.HTML(http.StatusOK, "pay.tmpl", gin.H{"channelIdList": channelIdList})
		})
	}
}
