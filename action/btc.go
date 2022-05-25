package action

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"wallet-srv/lib"
	"wallet-srv/lib/resp"
	"wallet-srv/lib/tx/btc"
	types "wallet-srv/lib/types"
	"wallet-srv/lib/wallet"
	"wallet-srv/model"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func BtcSignAction(c *gin.Context) {
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
		   }
		   **/
	var reqData types.BtcTx
	err := lib.DecodeHTTPReqJSON(c, &reqData)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("EncryptAction: ", err)
		return
	}

	paramChain := wallet.GetChainParam(wallet.SymbolBtc)
	changeAddress, err := btcutil.DecodeAddress(reqData.ChangeAddress, paramChain)

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
		glog.Error("BtcSignAction: ", err)
		return
	}

	//获取签名私钥
	w, err := getBtcWallet(txvinAddress)
	if err != nil {
		resp.Error(c, resp.SIGN_ERROR)
		glog.Error("BtcSignAction: ", err, " ", txvinAddress)
		return
	}

	tx, err := btc.NewBtcTransaction(unspent, txVout, changeAddress, feePerKb, paramChain)
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("BtcSignAction: ", err)
		return
	}

	err = tx.Sign(w)
	if err != nil {
		resp.Error(c, resp.SIGN_ERROR)
		glog.Error("BtcSignAction: ", err)
		return
	}

	buff := bytes.NewBuffer(make([]byte, 0, tx.Tx.SerializeSize()))
	tx.Tx.Serialize(buff)
	signTxHex := hex.EncodeToString(buff.Bytes())

	ret := tx.Decode()
	b, _ := json.MarshalIndent(ret, "", " ")
	fmt.Println("decoded tx:", string(b))

	glog.Info("BtcSignTx: ", signTxHex)
	resp.Ok(c, signTxHex)
}

func getBtcWallet(address string) (*wallet.BtcWallet, error) {

	cModel, table := wallet.GetDBModel("btc")
	var addrM = &model.CoinAddress{}
	addrM.Address = address
	cModel.GetData(table, addrM)

	if addrM.Address == "" {
		return nil, fmt.Errorf("No such address in the database.")
	}

	path, _ := wallet.MakeBip44Path(wallet.BTC, 0, 0, 0)
	s, _ := hex.DecodeString(addrM.PrivKey)
	paramChain := wallet.GetChainParam(wallet.SymbolBtc)
	w, err := wallet.NewBtcWalletByPath(wallet.SymbolBtc, path, s, wallet.SegWitNone, paramChain)
	if err != nil {
		return nil, err
	}

	return w, nil
}
