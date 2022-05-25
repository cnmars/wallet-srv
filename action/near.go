package action

import (
	"encoding/base64"
	"encoding/hex"
	"errors"

	"wallet-srv/lib"

	"wallet-srv/lib/pkg/near/serialize"
	"wallet-srv/lib/pkg/near/transaction"
	"wallet-srv/lib/resp"
	"wallet-srv/lib/wallet"

	ctypes "wallet-srv/lib/types"
	"wallet-srv/model"

	"wallet-srv/lib/decimal"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func NearSignAction(c *gin.Context) {
	/*
		{
			"from": "",
			"to": "",
			"amount": 1.231131,
			"last_blockhash": "" //最新的块hash
			"nonce": 0
		}
		**/
	var reqData ctypes.NearTx
	err := lib.DecodeHTTPReqJSON(c, &reqData)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("NearSignAction: DecodeHTTPReqJSON error", err)
		return
	}
	amount := decimal.NewFromFloat(reqData.Amount).Shift(24)
	fromAddress := reqData.From
	toAddress := reqData.To

	//获取私钥
	w, err := getNearPrivateKey(fromAddress)
	if err != nil {
		resp.Error(c, resp.FROM_ADDRESS)
		glog.Error("NearSignAction getSolPrivateKey error ", err)
		return
	}

	//构建TX交易
	tx, err := transaction.CreateTransaction(fromAddress, toAddress, w.DerivePublicKey(), reqData.LastBlockHash, reqData.Nonce)
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("NearSignAction transaction.CreateTransaction error ", err)
		return
	}

	ta, err := serialize.CreateTransfer(amount.String())
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("NearSignAction serialize.CreateTransfer error ", err)
		return
	}

	tx.SetAction(ta)
	txData, err := tx.Serialize()
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("NearSignAction tx.Serialize error ", err)
		return
	}

	sig, err := transaction.SignTransaction(hex.EncodeToString(txData), w.DerivePrivateKey())
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("NearSignAction transaction.SignTransaction error ", err)
		return
	}

	stx, err := transaction.CreateSignatureTransaction(tx, sig)
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("NearSignAction transaction.CreateSignatureTransaction error ", err)
		return
	}

	stxData, err := stx.Serialize()
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("NearSignAction stx.Serialize() error ", err)
		return
	}
	b64Data := base64.StdEncoding.EncodeToString(stxData)

	glog.Info("NearSignTx: ", b64Data)
	resp.Ok(c, b64Data)
}

func getNearPrivateKey(address string) (*wallet.NearWallet, error) {

	cModel, table := wallet.GetDBModel("near")
	var addrM = &model.CoinAddress{}
	addrM.Address = address
	cModel.GetData(table, addrM)
	if addrM.PrivKey == "" {
		return nil, errors.New("No such address in the database.")
	}

	s, _ := hex.DecodeString(addrM.PrivKey)
	w, err := wallet.NewNearWallet(s)
	if err != nil {
		return nil, err
	}

	return w, nil
}
