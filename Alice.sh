source ~/sgxsdk/environment
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client
go run main.go -port=3001 -grpc_port=50001 -peer_file_directory=data/peer/peer0.json -key_file=./data/key/k0 -channel_file=./data/channel/c0
