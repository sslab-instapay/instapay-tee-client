package config

import (
	"github.com/sslab-instapay/instapay-tee-client/model"
	"strconv"
	"os"
)

var AccountConfig = model.Account{}

func GetAccountConfig() model.Account {

	if AccountConfig.PublicKeyAddress != "" {
		return AccountConfig
	}

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
	}else if port == 3004{
		return model.Account{
			PublicKeyAddress: "0x0ED5B1a31b7D7CA0663f5298461FD44249475dde",
			PrivateKey: "8a0d44f5799a7af5bb375f5db8ba0d803775784b97fe542f1e7890bb9db881b7",
		}
	}else if port == 3005{
		return model.Account{
			PublicKeyAddress: "0xbDb2e1dfd4493C966F8c0289502E1674825B7D14",
			PrivateKey: "bdc7177b984a12e69fa25a478d3b771371986e1792f3fa662f095545513913f1",
		}
	}else if port == 3006{
		return model.Account{
			PublicKeyAddress: "0x15a3E2348bd529c6dF76C5f840714D1f05Ab6954",
			PrivateKey: "89bfcbdd969f87ed906a4417ec2b75d8a03d7a0bd720ec291991f1c308c70136",
		}
	}else if port == 3007{
		return model.Account{
			PublicKeyAddress: "0x2cc07D9903D3F15F5F8C343B34BB41EEB16c90FD",
			PrivateKey: "1e0487f9f86f2befdd5fd937be7f5d0c719c345b3f2e3c5b5bde2affe76f1c5d",
		}
	}else if port == 3008{
		return model.Account{
			PublicKeyAddress: "0x421d8797AE9d9AF38Aa7804ecCd14E986beDFb51",
			PrivateKey: "9138d38da7c99ec4478dc40be20e8fa060bae6cb640fd8bf7059605565213684",
		}
	}else if port == 3009{
		return model.Account{
			PublicKeyAddress: "0xCe61Cf7937B3f2EC01400C24ef6888F969685F24",
			PrivateKey: "1c58b2a767a1f3aeac81234009f5ad02e3d81d1b661610450b9ce9edd1f316dd",
		}
	}else if port == 3010{
		return model.Account{
			PublicKeyAddress: "0xb3c8B15925CE8E6fE018b68d18D69C59D33f5a2e",
			PrivateKey: "2c69eb06706903ac0698d8e3e7b4ed5115c155a4bcad5902c4e46004f973ddfa",
		}
	}
	return model.Account{
		PublicKeyAddress: "0xD03A2CC08755eC7D75887f0997195654b928893e",
		PrivateKey: "e113ff405699b7779fbe278ee237f2988b1e6769d586d8803860d49f28359fbd",
	}
}
