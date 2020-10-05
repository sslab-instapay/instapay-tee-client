package config

var EthereumConfig = map[string]string{
	/* web3 and ethereum */
	"wsHost":           "141.223.121.139",
	"wsPort":           "8881",
	"contractAddr":     "0x092d70BB5c1954F5Fa3EBbb282d0416a5e46c818",
	"contractSrcPath":  "../contracts/InstaPay.sol",
	"contractInstance": "",
	"web3":             "",
	"event":            "",

	/* grpc configuration */
	"serverGrpcHost": "141.223.121.139",
	"serverGrpcPort": "50004",
	"serverProto":    "",
	"server":         "",
	"myGrpcPort":     "", //process.argv[3]
	"clientProto":    "",
	"receiver":       "",
}
