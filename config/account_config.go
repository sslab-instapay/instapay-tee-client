package config

import (
	"github.com/sslab-instapay/instapay-tee-client/model"
)

var personalAccount = model.Account{}

func SetAccountConfig(publicKey string) model.Account {
	personalAccount.PublicKeyAddress = publicKey
	return personalAccount
}

func GetAccountConfig() model.Account {
	return personalAccount
}
