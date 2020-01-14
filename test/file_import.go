package main

import (
	"github.com/sslab-instapay/instapay-tee-client/util"
	"log"
)

func main(){
	util.SetPeerInformation("../data/peer/peer.json")

	_, _, err := util.GetPeerInformationByAddress("0xq2eqer2qeqwtqwet")
	if err != nil{
		log.Fatal(err)
	}

	util.AddOrUpdatePeerInformation("0x4451eerjg9o2eqwe", "141.223.423.21", 50001)

	util.ExportInformationToFile("../data/peer/peer.json")

}
