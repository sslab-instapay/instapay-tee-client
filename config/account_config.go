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
	// TODO: 이제 포트별 주소지정 필요 없이, SetAccountConfig에서 설정한 계정 정보 불러오기
	// port, _ := strconv.Atoi(os.Getenv("port"))
	// if port == 3001{
	// 	return model.Account{
	// 		PublicKeyAddress: "0xD03A2CC08755eC7D75887f0997195654b928893e",
	// 	}
	// }
	return personalAccount
}
