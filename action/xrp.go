package action

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strconv"
	"wallet-srv/lib"
	"wallet-srv/lib/resp"
	"wallet-srv/lib/tx/xrp"
	ctypes "wallet-srv/lib/types"
	"wallet-srv/lib/wallet"
	"wallet-srv/model"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func XrpSignAction(c *gin.Context) {
	/*
		   {
			   "from": "",
			   "to": "",
			   "amount": "121.210231",
			   "fee": "0.000012",
			   "assetcode": "", //
			   "assetissuer": "",
			   "sequence": 12,
		   }
		   **/
	var reqData ctypes.XrpTx
	err := lib.DecodeHTTPReqJSON(c, &reqData)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("XrpSignAction: DecodeHTTPReqJSON error", err)
		return
	}

	w, err := getXrpPrivateKey(reqData.From)
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("XrpSignAction: getData error ", err)
		return
	}

	tx, err := xrp.NewTxSign(
		w.DerivePrivateKey(),
		reqData.From,
		reqData.To,
		reqData.Amount,
		reqData.AssetCode,
		reqData.AssetIssuer,
		feeToInt64(reqData.Fee),
		&reqData.FromSeq,
	)
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("XrpSignAction: CreateTransaction error ", err)
		return
	}

	txid := tx.Hash.String()
	signTxHex := tx.TxnSignature.String()
	glog.Infof("XrpSignTx: txid=%s, signTx=%s", txid, signTxHex)

	var ret = &ctypes.Ret{
		TxId:   txid,
		SignTx: signTxHex,
	}

	resp.Ok(c, ret)
}

func getXrpPrivateKey(address string) (*wallet.XrpWallet, error) {
	cModel, table := wallet.GetDBModel("xrp")
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
	w, err := wallet.NewXrpWallet(s)
	if err != nil {
		return nil, fmt.Errorf("NewXrpWallet fail by privKey")
	}
	return w, nil
}

func feeToInt64(fee string) int64 {
	f, _ := strconv.ParseFloat(fee, 64)
	return int64(f * math.Pow10(6))
}
