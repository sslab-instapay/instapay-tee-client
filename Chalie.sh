source ~/sgxsdk/environment
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client/
go run main.go -port=3003 -grpc_port=50003 -peer_file_directory=data2/peer/peer.json