package action

import (
	"encoding/base64"
	"fmt"
	"wallet-srv/lib"
	"wallet-srv/lib/resp"
	"wallet-srv/lib/tx/avax"
	types "wallet-srv/lib/types"
	"wallet-srv/lib/wallet"
	"wallet-srv/model"

	"github.com/ava-labs/avalanchego/utils/crypto"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func AvaxSignAction(c *gin.Context) {
	/*
		   {

		   	"vin":[
				   {"txid":"xxxxxxx", "vout":0, "amount": 11, "address": "xxxx", "asset_id": "xxxxx"},
				   {"txid":"xxxxxxx", "vout":1, "amount": 111, "address": "xxxx", "asset_id": "xxxxx"},
				   {"txid":"xxxxxxx", "vout":2, "amount": 11, "address": "xxxx", "asset_id": "xxxxx"},
			   ],
		   	"vout":[
				   {"address": "1xxx99a9axxx", "amount": 1212, "asset_id": "xxxxx"},
				   {"address": "1xxx99a9axxx", "amount": 11, "asset_id": "xxxxx"},
				   {"address": "1xxx99a9axxx", "amount": 11, "asset_id": "xxxxx"},
			   ],
			"block_chainid": "xxxxxxxxxxxxx",//
			"chain_id": 1,
			"memo": ""
		   }
		   **/
	var reqData types.AvaxTx
	err := lib.DecodeHTTPReqJSON(c, &reqData)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("EncryptAction: ", err)
		return
	}

	baseTx := avax.NewTransaction()
	err = baseTx.AddInput(reqData.Inputs)
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("AvaxSignAction: ", err)
		return
	}

	err = baseTx.AddOutput(reqData.Outputs)
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("AvaxSignAction: ", err)
		return
	}

	//设置chainId和blockChainId
	baseTx.SetChainId(reqData.ChainId)
	baseTx.SetBlockChainId(reqData.BlockChainId)

	baseTx.Build()
	var signer []*crypto.PrivateKeySECP256K1R

	for _, in := range reqData.Inputs {
		privkey, err := getAvaxWallet(in.Address)
		if err != nil {
			continue
		}
		signer = append(signer, privkey.(*crypto.PrivateKeySECP256K1R))
	}

	var signOwner [][]*crypto.PrivateKeySECP256K1R
	signOwner = append(signOwner, signer)

	if err := baseTx.Sign(signOwner); err != nil {
		resp.Error(c, resp.SIGN_ERROR)
		glog.Error("AvaxSignAction: ", err)
		return
	}
	txid := baseTx.Tx.ID().String()
	signTxHex := base64.StdEncoding.EncodeToString(baseTx.Tx.Bytes())

	ret := &types.Ret{
		TxId:   txid,
		SignTx: signTxHex,
	}

	glog.Info("Avax TxId: ", txid)
	glog.Info("AvaxSignTx: ", signTxHex, txid)
	resp.Ok(c, ret)
}

func getAvaxWallet(address string) (crypto.PrivateKey, error) {

	cModel, table := wallet.GetDBModel("avax")
	var addrM = &model.CoinAddress{}
	addrM.Address = address
	cModel.GetData(table, addrM)

	if addrM.Address == "" {
		return nil, fmt.Errorf("No such address in the database.")
	}
	factory := crypto.FactorySECP256K1R{}
	privKey, err := factory.ToPrivateKey([]byte(addrM.PrivKey))
	if err != nil {
		return nil, fmt.Errorf("ToPrivateKey error: %v", err)
	}

	return privKey, nil
}
