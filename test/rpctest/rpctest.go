package rpctest

import (
	"encoding/json"
	"fmt"
	"time"
	"wallet-srv/lib"
	"wallet-srv/lib/rpc"
	"wallet-srv/lib/types"

	"github.com/golang/glog"
)

const SIGN_HTTP_API = "http://127.0.0.1:4040/v1"

type RespData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type RespSignHex struct {
	Content string
}

func NewClient() *rpc.RpcClient {

	return rpc.New(SIGN_HTTP_API, "", "")
}

func DoSign(client *rpc.RpcClient, method string, req interface{}, signHex *RespSignHex) error {

	url := fmt.Sprintf("%s%s", SIGN_HTTP_API, method)

	var res *RespData

	reqJson, err := json.Marshal(req)
	if err != nil {
		return err
	}

	glog.Info("reqJson->", string(reqJson))
	reqContent, err := lib.APIAESEncrypt(string(reqJson))
	if err != nil {
		return err
	}
	reqToken, _ := lib.APIAESEncrypt(fmt.Sprintf("%d", time.Now().Unix()))
	Req := &types.ReqBody{
		Content: reqContent,
		Token:   reqToken,
	}

	Res := &types.RespBody{}
	err = client.PostRequest(url, Req, Res)
	if err != nil {
		glog.Error("PostRequest: ", url, " req->", req, " res->", res)
		return err
	}

	glog.Info(Res)

	if Res.Code != 0 {
		return fmt.Errorf("%v", Res.Message)
	}

	dstr, err := lib.APIAESDecrypt(Res.Data)
	if err != nil {
		return err
	}

	signHex.Content = dstr

	return nil
}
