source ~/sgxsdk/environment
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client/
go run main.go -port=3002 -grpc_port=50002 -peer_file_directory=data/peer/peer.json -key_file=./data/key/k1