package controller

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/sslab-instapay/instapay-tee-client/config"
	"github.com/sslab-instapay/instapay-tee-client/service"
	"github.com/sslab-instapay/instapay-tee-client/repository"
	"log"
	"reflect"
	"unsafe"
	"github.com/sslab-instapay/instapay-tee-client/util"
	"io/ioutil"
	"os"
	"fmt"
)

func AccountInformationHandler(context *gin.Context) {
	account := config.GetAccountConfig()
	balance, _ := service.GetBalance()

	lockedBalance, err := repository.GetAllChannelsLockedBalance()
	if err != nil {
		log.Print(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		convertedBalance, _ := balance.Int64()
		totalBalance := convertedBalance - lockedBalance
		context.JSON(http.StatusOK, gin.H{"address": account.PublicKeyAddress, "balance": totalBalance})
	}

}

func CreateAccountHandler(context *gin.Context) {

	var password = context.Param("password")
	var sig1 *C.uchar = C.ecall_create_account_w()
	hdr1 := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(sig1)),
		Len:  20,
		Cap:  20,
	}

	s1 := *(*[]C.uchar)(unsafe.Pointer(&hdr1))
	for i := C.uint(0); i < 20; i++ {
		fmt.Printf("%02x", s1[i])
	}

	encryptedData := util.Encrypt([]byte(s1), password)

	err := ioutil.WriteFile("test.txt", encryptedData, os.FileMode(644))
	if err != nil {
		log.Println(err)
	}

	//TODO TEE에서 데이터 실링을 통해 개인키 파일 암호화.

	config.AccountConfig.PublicKeyAddress = ""
	config.AccountConfig.PrivateKey = C.CString(sig1)

	log.Println("--- Create SuccessFully ---")
}
