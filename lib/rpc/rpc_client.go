package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

type RpcClient struct {
	rpcUrl      string
	rpcUser     string
	rpcPassword string
}

type RequestBody struct {
	ReqNotHaveParams
	Params []interface{} `json:"params"`
}
type ReqNotHaveParams struct {
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Id      int    `json:"id"`
}
type RespBody struct {
	JsonRpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   interface{} `json:"error"`
	Id      int         `json:"id"`
}

//初始化一个rpc客户端
func New(url, user, password string) *RpcClient {
	return &RpcClient{
		rpcUrl:      url,
		rpcUser:     user,
		rpcPassword: password,
	}
}

func (rpc *RpcClient) SendRequest(method string, params []interface{}) ([]byte, error) {
	id := rand.Intn(10000)
	var (
		reqBytes []byte
		err      error
	)
	if params != nil {
		var reqBody RequestBody
		reqBody.JsonRpc = "2.0"
		reqBody.Id = id
		reqBody.Method = method
		reqBody.Params = params
		reqBytes, err = json.Marshal(reqBody)
	} else {
		var reqBody ReqNotHaveParams
		reqBody.JsonRpc = "2.0"
		reqBody.Id = id
		reqBody.Method = method
		reqBytes, err = json.Marshal(reqBody)
	}
	if err != nil {
		return nil, err
	}
	reqBuf := bytes.NewBuffer(reqBytes)
	var (
		req *http.Request
	)

	if req, err = http.NewRequest(http.MethodPost, rpc.rpcUrl, reqBuf); err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	//设置rpc的用户和密码
	//如果为空就不设置
	if rpc.rpcUser != "" && rpc.rpcPassword != "" {
		req.SetBasicAuth(rpc.rpcUser, rpc.rpcPassword)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resp, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	//解析resp
	var response RespBody
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, err
	}

	if response.Error != nil {
		rpcErr := response.Error.(map[string]interface{})
		return nil, fmt.Errorf("rpc get error,Code=【%d】,Message=【%s】", int(rpcErr["code"].(float64)), rpcErr["message"].(string))
	}

	//如果返回的结果直接是一个string，就不在做json处理了，直接返回
	switch response.Result.(type) {
	case string:
		return []byte(response.Result.(string)), nil
	case float64:
		f := strconv.FormatFloat(response.Result.(float64), 'f', -1, 64)
		return []byte(f), nil
	default:
		data, err := json.Marshal(response.Result)
		if err != nil {
			return nil, fmt.Errorf("marshal result error,Err=【%v】", err)
		}
		return data, nil
	}
}

func (rpc *RpcClient) PostRequest(url string, reqBody interface{}, resBody interface{}) error {
	rpc.rpcUrl = url

	var (
		reqBytes []byte
		err      error
	)

	reqBytes, err = json.Marshal(reqBody)
	if err != nil {
		return err
	}
	reqBuf := bytes.NewBuffer(reqBytes)
	var (
		req *http.Request
	)

	if req, err = http.NewRequest(http.MethodPost, rpc.rpcUrl, reqBuf); err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	//设置rpc的用户和密码
	//如果为空就不设置
	if rpc.rpcUser != "" && rpc.rpcPassword != "" {
		req.SetBasicAuth(rpc.rpcUser, rpc.rpcPassword)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	//解析resp
	if err := json.Unmarshal(resp, &resBody); err != nil {
		return fmt.Errorf("parse resp error,Err=【%v】", err)
	}

	return nil
}
