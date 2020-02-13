package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"github.com/sslab-instapay/instapay-tee-client/repository"
	"github.com/sslab-instapay/instapay-tee-client/config"
	"github.com/sslab-instapay/instapay-tee-client/service"
)

func RegisterViewRouter(router *gin.Engine) {

	viewRouter := router.Group("templates")
	{
		// account 리스트
		viewRouter.GET("accounts/list", func(context *gin.Context) {
			//gin.H 부분에서 변수 다루는 것.
			account := config.GetAccountConfig()
			balance, _ := service.GetBalance()
			balanceStr := balance.String()
			context.HTML(http.StatusOK, "account.tmpl", gin.H{
				"account": account,
				"balance": balanceStr,
			})
		})

		// channel 리스트
		viewRouter.GET("channels/list", func(context *gin.Context) {
			channelList, err := repository.GetOpenedChannelList()
			if err != nil {
				log.Println(err)
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

		viewRouter.GET("channels/pay", func(context *gin.Context) {
			openedChannelList, err := repository.GetOpenedChannelList()
			if err != nil {
				log.Println(err)
			}

			var idList []int
			for _, channel := range openedChannelList{
				idList = append(idList, int(channel.ChannelId))
			}

			context.HTML(http.StatusOK, "pay.tmpl", gin.H{"channelIdList": idList})
		})
	}
}
