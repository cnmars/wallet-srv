package action

import (
	"encoding/hex"
	"errors"
	"fmt"
	"wallet-srv/lib"
	"wallet-srv/lib/filutil"
	"wallet-srv/lib/resp"
	"wallet-srv/lib/tx/fil"
	ctypes "wallet-srv/lib/types"
	"wallet-srv/lib/wallet"
	"wallet-srv/model"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func FilSignAction(c *gin.Context) {
	/*
		   {
			   "from": "",
			   "to": "",
			   "amount": 12313,		//单位为Wei
			   "nonce": 0,
			   "gaslimit": 60000,
			   "gasfee": 2121231, //单位为Wei
			   "gaspremium": 10000,
			   "param": ""
		   }
		   **/
	var reqData ctypes.FilTx
	err := lib.DecodeHTTPReqJSON(c, &reqData)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("FilSignAction: DecodeHTTPReqJSON error", err)
		return
	}

	nonce := uint64(reqData.Nonce)
	value := reqData.Amount
	fromAddress, err := filutil.NewFromString(reqData.From)
	if err != nil {
		resp.Error(c, resp.FROM_ADDRESS)
		glog.Error("FilSignAction: param error", err)
		return
	}
	toAddress, err := filutil.NewFromString(reqData.To)
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("FilSignAction: to address error", err)
		return
	}

	gasLimit := int64(reqData.GasLimit)
	gasFee := reqData.GasFee
	gasPremium := reqData.GasPremium
	param := []byte(reqData.Param)

	w, err := getFilPrivateKey(reqData.From)
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("FilSignAction: getData error ", err)
		return
	}

	tx, err := fil.CreateTransaction(
		fromAddress,
		toAddress,
		value,
		gasLimit,
		gasFee,
		gasPremium,
		nonce,
		0,
		param,
	)
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("FilSignAction: CreateTransaction error ", err)
		return
	}

	signTxRaw, err := fil.SignMessage(w.GetPrivKey(), tx)
	if err != nil {
		resp.Error(c, resp.SIGN_ERROR)
		glog.Error("FilSignAction: SignMessage error ", err)
		return
	}

	if !fil.VerifyMessage(signTxRaw) {
		resp.Error(c, resp.SIGN_ERROR)
		glog.Error("FilSignAction: sign verify error ")
		return
	}

	signTxHex := hex.EncodeToString(signTxRaw.Signature.Data)
	txid := signTxRaw.Cid().String()

	glog.Infof("FilSignTx: txid=%s, signTx=%s", txid, signTxHex)

	var ret = &ctypes.Ret{
		TxId:   txid,
		SignTx: signTxHex,
	}

	resp.Ok(c, ret)
}

func getFilPrivateKey(address string) (*wallet.FilWallet, error) {
	cModel, table := wallet.GetDBModel("fil")
	var addrM = &model.CoinAddress{}
	addrM.Address = address
	cModel.GetData(table, addrM)
	if addrM.PrivKey == "" {
		return nil, errors.New("No such address in the database.")
	}

	s, err := hex.DecodeString(addrM.PrivKey)
	if err != nil {
		return nil, fmt.Errorf("address %s privKey error", address)
	}
	w, err := wallet.NewFilWallet(s)
	if err != nil {
		return nil, fmt.Errorf("NewFilWallet fail by privKey")
	}
	return w, nil
}
