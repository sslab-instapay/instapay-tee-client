source ~/sgxsdk/environment
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client/
go run main_Carol.go -port=3003 -grpc_port=50003 -peer_file_directory=data/peer/peer2.json -key_file=./data/key/k2 -channel_file=./data/channel/c2
