package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
	"wallet-srv/conf"
	"wallet-srv/lib"
	"wallet-srv/lib/rpc"
	wtypes "wallet-srv/lib/types"

	"github.com/golang/glog"
)

//encrypt privkey
func DBEncrypt(addr *CoinAddress) error {

	//first encrypt seed
	var err error
	var privKey, xPrivKey string
	addr.Multi = 1
	privKey, err = lib.SignAESEncrypt(addr.PrivKey)
	if err != nil {
		glog.Info("AESEncrypt seed: ", addr.PrivKey, " fail.")
		return errors.New("AESEncrypt seed fail. ")
	}
	addr.PrivKey = privKey

	//second encrypt seed
	//xPrivKey, err = lib.XXTEAEncrypt(privKey)
	xPrivKey, err = encryptAPI(privKey, "/encrypt")
	if err == nil {
		addr.PrivKey = xPrivKey
		addr.Multi = 2
	}

	return nil

}

//decrypt db privkey
func DBDecrypt(addr *CoinAddress) error {
	var err error
	var privKey string

	privKey = addr.PrivKey

	//if secord encrypt
	if addr.Multi == 2 {
		//privKey, err = lib.XXTEADecrypt(privKey)
		privKey, err = encryptAPI(privKey, "/decrypt")
		if err != nil {
			glog.Fatal("the second XXTEADecrypt decrypt fail.")
			return errors.New("the second XXTEADecrypt decrypt fail. ")
		}
	}
	addr.PrivKey = privKey

	//decrypt the first encrypt privkey
	privKey, err = lib.SignAESDecrypt(privKey)
	if err != nil {
		glog.Fatal("the first SignAESDecrypt decrypt fail.")
		return errors.New("the first SignAESDecrypt decrypt fail.")
	}

	addr.PrivKey = privKey

	return nil
}

//encrypt privkey by encrypt api service
func encryptAPI(body string, method string) (string, error) {

	var d wtypes.Content
	d.Data = body
	dJson, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	eBody, err := lib.APIAESEncrypt(string(dJson))
	if err != nil {
		return "", err
	}
	eTime, _ := lib.APIAESEncrypt(fmt.Sprintf("%d", time.Now().Unix()))

	var reqData wtypes.ReqBody
	reqData.Content = eBody
	reqData.Token = eTime

	url := fmt.Sprintf("%s%s", conf.ENCRYPT_HTTP_API, method)

	client := rpc.New(url, "", "")

	var resBody *wtypes.RespBody
	err = client.PostRequest(url, reqData, &resBody)
	if err != nil || resBody.Code != 0 {
		glog.Error("secondEncrypt error: ", resBody.Message)
		return "", err
	}

	ePrivKey, err := lib.APIAESDecrypt(resBody.Data)
	if err != nil {
		return "", err
	}

	return strings.Trim(ePrivKey, "\""), nil
}

func Encode(str string) (string, error) {
	return encryptAPI(str, "/encrypt")
}

func Decode(str string) (string, error) {
	return encryptAPI(str, "/decrypt")
}
