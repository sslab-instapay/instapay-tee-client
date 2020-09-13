# Go Client for InstaPay

인스타페이 프로토콜을 위한 클라이언트 프로그램

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



## 프로그램은 크게 세 가지 파트로 구성

1. 이더리움 결제채널 이벤트 받는 쓰레드.
2. Grpc 서버 (다른 사용자와 통신할 때 사용)
3. 웹 서버 (사용자 정보 확인 결제 채널 정보 및 결제 요청 인터페이스)


## 디렉토리 별 설명
1. config => 이더리움, 계정 정보 관련 셋업
2. router => 사용자 인터페이스 관련 웹 경로 정보
ex) templates/channels/list
3. controller => 컨트롤러 로직(웹 관련)
4. db => 안쓰는 것.
5. grpc => 결제채널 당사자끼리 결제할 때 메시지 교환은 grpc를 통해 
6. model => 결제채널, 결제 데이터 등 데이터 형식 정의
7. service => 이더리움 관련 이벤트 리스너 
8. templates => 사용자 웹 관련 템플릿파일
9. peers => 피어 정보 로드 (for 데모)
10. data => 프리 로드할 sealed된 공개/키 페어 파일


## 실행 스크립트 예

1. Alice.sh => 사용자 앨리스 가정
```sh
source ~/sgxsdk/environment
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/home/xiaofo/instapay/src/github.com/sslab-instapay/instapay-tee-client
go run main.go -port=3001 -grpc_port=50001 -peer_file_directory=data/peer/peer0.json -key_file=./data/key/k0 -channel_file=./data/channel/c0
```

