package util

import (
	"testing"
	"github.com/sslab-instapay/instapay-tee-client/util"
		"os"
	"io/ioutil"
	"log"
	"github.com/magiconair/properties/assert"
	"fmt"
)

func TestEncryptAndDecrypt(t *testing.T){

	testString := "qwetqwet"

	encryptedData := util.Encrypt([]byte(testString), "password")

	err := ioutil.WriteFile("test.txt", encryptedData, os.FileMode(644))

	if err != nil{
		log.Println(err)
	}

	data, err := ioutil.ReadFile("test.txt")
	if err != nil{
		log.Println(err)
	}

	decryptedData := util.Decrypt(data, "password")
	fmt.Println(string(decryptedData))

	assert.Equal(t, testString, string(decryptedData))

}