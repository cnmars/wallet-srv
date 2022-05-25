package action

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"wallet-srv/lib"
	"wallet-srv/lib/resp"
	"wallet-srv/lib/tx/trx"
	ctypes "wallet-srv/lib/types"
	"wallet-srv/lib/wallet"
	"wallet-srv/model"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

//TRX验签
func TrxSignAction(c *gin.Context) {
	/*
		   {
			   "from": "",
			   "raw_data_hex": "" //请求端通过调用 /wallet/createtransaction 或 /wallet/triggersmartcontract返回值里面的raw_data_hex值
		   }
		   见 https://tronprotocol.github.io/documentation-zh/api/http/
		   **/
	var reqData ctypes.TrxTx
	err := lib.DecodeHTTPReqJSON(c, &reqData)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("TrxSignAction DecodeHTTPReqJSON error", err)
		return
	}

	fromAddress := reqData.From
	rawData, _ := hex.DecodeString(reqData.RawDataHex)

	if !wallet.CheckTrxAddress(fromAddress) {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("TrxSignAction ReqData address error")
		return
	}

	privKey, err := getTrxPrivateKey(fromAddress)
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("TrxSignAction getData error ", err)
		return
	}

	//签名交易
	signTx, err := trx.SignTx(privKey, rawData)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("TrxSignAction SignTx error: ", err)
		return
	}
	signHex := hex.EncodeToString(signTx)
	glog.Info("TrxSignTx: ", signHex)
	resp.Ok(c, signHex)
}

func getTrxPrivateKey(address string) (*ecdsa.PrivateKey, error) {
	cModel, table := wallet.GetDBModel("trx")
	var addrM = &model.CoinAddress{}
	addrM.Address = address
	cModel.GetData(table, addrM)
	if addrM.PrivKey == "" {
		return nil, errors.New("No such address in the database.")
	}

	s, _ := hex.DecodeString(addrM.PrivKey)
	path, _ := wallet.MakeBip44Path(wallet.TRX, 0, 0, 0)
	w, err := wallet.NewEthWalletByPath(path, s, wallet.EthChainId)
	if err != nil {
		return nil, err
	}

	return w.DeriveNativePrivateKey(), nil
}
