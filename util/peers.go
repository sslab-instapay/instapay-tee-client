package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/sslab-instapay/instapay-tee-client/model"
)

var peerInformations = model.PeerInformations{}

func SetPeerInformation(directory string) {
	log.Println(directory)
	file, err := ioutil.ReadFile(directory)
	if err != nil {
		log.Println("Read File Error")
	}

	data := model.PeerInformations{}

	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < len(data.PeerInformationList); i++ {
		fmt.Println("Peer Address: ", &data.PeerInformationList[i])
		fmt.Println("Ip Address: ", data.PeerInformationList[i].IpAddress)
		fmt.Println("GRPC: ", data.PeerInformationList[i].GrpcPort)
	}

	peerInformations.PeerInformationList = data.PeerInformationList
}

func GetPeerInformationByAddress(publicKeyAddress string) (model.PeerInformation, int, error) {
	log.Println("Peer Address : ", publicKeyAddress)
	for i := 0; i < len(peerInformations.PeerInformationList); i++ {
		if peerInformations.PeerInformationList[i].PublicKeyAddress == publicKeyAddress {
			return peerInformations.PeerInformationList[i], i, nil
		}
	}
	return model.PeerInformation{}, -1, errors.New("There is no peer information")
}

func ExportInformationToFile(directory string) {
	file, _ := json.MarshalIndent(peerInformations, "", " ")

	err := ioutil.WriteFile(directory, file, 0644)
	if err != nil {
		log.Println(err)
	}
}

func AddOrUpdatePeerInformation(address string, ipAddress string, grpcPort int) {
	peer := model.PeerInformation{PublicKeyAddress: address, IpAddress: ipAddress, GrpcPort: grpcPort}

	_, idx, err := GetPeerInformationByAddress(address)
	if err != nil {
		peerInformations.PeerInformationList = append(peerInformations.PeerInformationList[:idx], peerInformations.PeerInformationList[idx+1:]...)
		peerInformations.PeerInformationList = append(peerInformations.PeerInformationList, peer)
	} else {
		peerInformations.PeerInformationList = append(peerInformations.PeerInformationList, peer)
	}
}
