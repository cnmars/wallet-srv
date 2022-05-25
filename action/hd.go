package action

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"wallet-srv/lib"
	"wallet-srv/lib/resp"
	"wallet-srv/lib/tx/btc"
	types "wallet-srv/lib/types"
	"wallet-srv/lib/wallet"
	"wallet-srv/model"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/txscript"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func HdSignAction(c *gin.Context) {
	/*
		   {

		   	"vin":[
				   {"txid":"xxxxxxx", "vout":0, "amount": 0.001, "address": "xxxx"},
				   {"txid":"xxxxxxx", "vout":1, "amount": 0.002, "address": "xxxx"},
				   {"txid":"xxxxxxx", "vout":2, "amount": 0.0003, "address": "xxxx"},
			   ],
		   	"vout":[
				   {"address": "1xxx99a9axxx", "amount": 0.0002},
				   {"address": "1xxx99a9axxx", "amount": 0.0001},
				   {"address": "1xxx99a9axxx", "amount": 0.001},
			   ],
			"change": "xxxxxxxxxxxxx",// 找零地址
			"feekb": 8000 //手续费每kb多少sat
			"coin": "bch"
		   }
		   **/
	var reqData types.HdTx
	err := lib.DecodeHTTPReqJSON(c, &reqData)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("HdSignAction: ", err)
		return
	}

	symbol := strings.ToUpper(reqData.Coin)
	paramChain := wallet.GetChainParam(symbol)
	changeAddress, _ := btcutil.DecodeAddress(reqData.ChangeAddress, paramChain)

	var txvinAddress string
	feePerKb := int64(reqData.FeePerKb)

	unspent := make([]btc.BtcUnspent, 0, len(reqData.TxVin))
	for _, in := range reqData.TxVin {
		addr, err := btcutil.DecodeAddress(in.Address, paramChain)
		if err != nil {
			continue
		}
		txvinAddress = in.Address

		pkScript, _ := txscript.PayToAddrScript(addr)
		utxo := btc.BtcUnspent{
			TxID:         in.Txid,
			Amount:       float64(in.Amount),
			Vout:         in.Vout,
			ScriptPubKey: hex.EncodeToString(pkScript),
			Address:      in.Address,
		}
		unspent = append(unspent, utxo)
	}

	txVout := make([]btc.BtcOutput, 0, len(reqData.TxVout))
	for _, vout := range reqData.TxVout {
		outAddress, err := btcutil.DecodeAddress(vout.Address, paramChain)
		if err != nil {
			continue
		}
		out := btc.BtcOutput{
			Address: outAddress,
			Amount:  btc.BtcToSatoshi(vout.Amount),
		}
		txVout = append(txVout, out)
	}

	if len(unspent) == 0 || len(txVout) == 0 {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("HdSignAction: ", err)
		return
	}

	//获取签名私钥
	w, err := getBtcWallet(txvinAddress)
	if err != nil {
		resp.Error(c, resp.SIGN_ERROR)
		glog.Error("HdSignAction: ", err, " ", txvinAddress)
		return
	}

	tx, err := btc.NewBtcTransaction(unspent, txVout, changeAddress, feePerKb, paramChain)
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("HdSignAction: ", err)
		return
	}

	err = tx.Sign(w)
	if err != nil {
		resp.Error(c, resp.SIGN_ERROR)
		glog.Error("HdSignAction: ", err)
		return
	}

	buff := bytes.NewBuffer(make([]byte, 0, tx.Tx.SerializeSize()))
	tx.Tx.Serialize(buff)
	signTxHex := hex.EncodeToString(buff.Bytes())

	ret := tx.Decode()
	b, _ := json.MarshalIndent(ret, "", " ")
	fmt.Println("decoded tx:", string(b))

	glog.Info("[", symbol, "]", "HdSignTx: ", signTxHex)
	resp.Ok(c, signTxHex)
}

func getHdWallet(symbol, address string) (*wallet.BtcWallet, error) {

	cModel, table, coinType := getModelChainId(symbol)
	var addrM = &model.CoinAddress{}
	addrM.Address = address
	cModel.GetData(table, addrM)

	if addrM.Address == "" {
		return nil, fmt.Errorf("No such address in the database.")
	}

	path, _ := wallet.MakeBip44Path(coinType, 0, 0, 0)
	s, _ := hex.DecodeString(addrM.PrivKey)
	paramChain := wallet.GetChainParam(symbol)
	w, err := wallet.NewBtcWalletByPath(symbol, path, s, wallet.SegWitNone, paramChain)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func getModelChainId(symbol string) (model.IfAddress, string, uint32) {
	m, table := wallet.GetDBModel(symbol)
	return m, table, wallet.GetCoinType(symbol)
}
