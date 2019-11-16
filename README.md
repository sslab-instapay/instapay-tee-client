# Go Client for InstaPay

## sgx environment
```sh
source $SGX_SDK/environment
```

## environment variable
```sh
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:$GOPATH/src/github.com/sslab-instapay/instapay-tee-client
```

## run
```sh
go run main.go -port=3001 -grpc_port=50001 -database_name=instapay-client
```