package config

import (
	"github.com/sslab-instapay/instapay-tee-client/model"
	"strconv"
	"os"
)

var personalAccount = model.Account{}

func SetAccountConfig(publicKey string) model.Account{
	personalAccount.PublicKeyAddress = publicKey
	return personalAccount
}

func GetAccountConfig() model.Account {
	port, _ := strconv.Atoi(os.Getenv("port"))
	if port == 3001{
		return model.Account{
			PublicKeyAddress: "0xD03A2CC08755eC7D75887f0997195654b928893e",
		}
	}
	return personalAccount
}
