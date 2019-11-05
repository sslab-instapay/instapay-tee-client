package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/sslab-instapay/instapay-go-client/config"
	"github.com/sslab-instapay/instapay-go-client/service"
	"github.com/sslab-instapay/instapay-go-client/repository"
	"log"
)

func AccountInformationHandler(context *gin.Context) {
	account := config.GetAccountConfig()
	balance := service.GetBalance()

	lockedBalance, err := repository.GetAllChannelsLockedBalance()
	if err != nil {
		log.Print(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		convertedBalance, _ := balance.Int64()
		totalBalance := convertedBalance - lockedBalance
		context.JSON(http.StatusOK, gin.H{"address": account.PublicKeyAddress, "balance": totalBalance})
	}

}
