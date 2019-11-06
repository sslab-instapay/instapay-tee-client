package config

import (
	"github.com/sslab-instapay/instapay-tee-client/model"
	"strconv"
	"os"
)

var AccountConfig = map[string]string{
	/* web3 and ethereum */
	"PublicKeyAddress": GetAccountConfig().PublicKeyAddress,
	"PrivateKey": GetAccountConfig().PrivateKey,
}


func GetAccountConfig() model.Account {
	port, _ := strconv.Atoi(os.Getenv("port"))
	if port == 3001{
		return model.Account{
			PublicKeyAddress: "0xD03A2CC08755eC7D75887f0997195654b928893e",
			PrivateKey: "e113ff405699b7779fbe278ee237f2988b1e6769d586d8803860d49f28359fbd",
		}
	}else if port == 3002{
		return model.Account{
			PublicKeyAddress: "0x0b4161ad4f49781a821C308D672E6c669139843C",
			PrivateKey: "240af81838ad22e8baa5c6223c7c7e112b091ba50e6fb396c0dc2b84cf034169",
		}

	}else if port == 3003{
		return model.Account{
			PublicKeyAddress: "0x78902c58006916201F65f52f7834e467877f0500",
			PrivateKey: "3038465f2b9be0048caa9f33e25b5dc50252f04c078aaddfbea74f26cdeb9f3c",
		}
		//TODO 개인키 공개키 발급 new
	}else if port == 3004{
		return model.Account{
			PublicKeyAddress: "0x78902c58006916201F65f52f7834e467877f0500",
			PrivateKey: "3038465f2b9be0048caa9f33e25b5dc50252f04c078aaddfbea74f26cdeb9f3c",
		}
	}else if port == 3005{
		return model.Account{
			PublicKeyAddress: "0x78902c58006916201F65f52f7834e467877f0500",
			PrivateKey: "3038465f2b9be0048caa9f33e25b5dc50252f04c078aaddfbea74f26cdeb9f3c",
		}
	}
	return model.Account{
		PublicKeyAddress: "0xD03A2CC08755eC7D75887f0997195654b928893e",
		PrivateKey: "e113ff405699b7779fbe278ee237f2988b1e6769d586d8803860d49f28359fbd",
	}
}
