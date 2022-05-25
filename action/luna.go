package action

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strconv"
	"wallet-srv/lib"
	"wallet-srv/lib/pkg/terra/msg"
	"wallet-srv/lib/pkg/terra/tx"
	"wallet-srv/lib/resp"
	"wallet-srv/lib/tx/luna"
	"wallet-srv/lib/wallet"
	"wallet-srv/model"

	ctypes "wallet-srv/lib/types"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func LunaSignAction(c *gin.Context) {
	/*
		   {
			   "from": "",
			   "to": "",
			   "amount": "121.210231",
			   "fee": "0.000012",
			   "gaslimit": 100,
			   "sequence": 12,
			   "coin": "" //默认为空的话为luna
		   }
		   **/
	var reqData ctypes.LunaTx
	err := lib.DecodeHTTPReqJSON(c, &reqData)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("LunaSignAction: DecodeHTTPReqJSON error", err)
		return
	}

	from, err := msg.AccAddressFromBech32(reqData.From)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("LunaSignAction: from AccAddressFromBech32 error", err)
		return
	}

	to, err := msg.AccAddressFromBech32(reqData.To)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("LunaSignAction: to AccAddressFromBech32 error", err)
		return
	}
	coin := "luna"
	if reqData.Coin != "" {
		coin = reqData.Coin
	}

	amount, err := strToInt64(reqData.Amount)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("LunaSignAction: amount strToInt64 error", err)
		return
	}
	feeAmount, err := strToInt64(reqData.FeeAmount)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("LunaSignAction: FeeAmount strToInt64 error", err)
		return
	}

	w, err := getLunaPrivateKey(reqData.From)
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("LunaSignAction: getData error ", err)
		return
	}

	txSign, err := luna.NewTxSign(luna.TxOptions{
		Msgs: []msg.Msg{
			msg.NewMsgSend(from, to, msg.NewCoins(msg.NewInt64Coin(coin, amount))), // 100UST
		},
		Memo:          "",
		SignMode:      tx.SignModeDirect,
		AccountNumber: 1,
		Sequence:      reqData.Seq,
		GasLimit:      reqData.GasLimit,
		FeeAmount:     msg.NewCoins(msg.NewInt64Coin(coin, feeAmount)),
	}, w.DeriveNativePrivateKey())
	if err != nil {
		resp.Error(c, resp.SIGN_ERROR)
		glog.Error("LunaSignAction: CreateTransaction error ", err)
		return
	}

	signTx, err := txSign.GetTxBytes()
	if err != nil {
		resp.Error(c, resp.SIGN_ERROR)
		glog.Error("LunaSignAction: txSign.GetTxBytes error ", err)
		return
	}
	signTxHex := hex.EncodeToString(signTx)
	glog.Infof("LunaSignTx:, signTx=%s", signTxHex)

	resp.Ok(c, signTxHex)
}

func getLunaPrivateKey(address string) (*wallet.LunaWallet, error) {

	cModel, table := wallet.GetDBModel("luna")
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
	w, err := wallet.NewLunaWallet(s)
	if err != nil {
		return nil, fmt.Errorf("NewXrpWallet fail by privKey")
	}
	return w, nil
}

func strToInt64(amount string) (int64, error) {
	f, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}
	return int64(f * math.Pow10(6)), nil
}
