package action

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"

	"wallet-srv/lib"

	"wallet-srv/lib/pkg/solana"
	"wallet-srv/lib/resp"
	"wallet-srv/lib/wallet"

	ctypes "wallet-srv/lib/types"
	"wallet-srv/model"

	"wallet-srv/lib/pkg/solana/base58"
	"wallet-srv/lib/pkg/solana/token"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func SolSignAction(c *gin.Context) {
	/*
		{
			"from": "",
			"to": "",
			"amount": 12313,
			"last_blockhash": "" //last block hash
			"token": "",
			"decimal": 9
		}
		**/
	var reqData ctypes.SolTx
	err := lib.DecodeHTTPReqJSON(c, &reqData)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("SolSignAction: DecodeHTTPReqJSON error", err)
		return
	}

	value := reqData.Amount
	fromAddress, err := base58.Decode(reqData.From)
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("SolSignAction ReqData from address error ")
		return
	}
	toAddress, err := base58.Decode(reqData.To)
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("SolSignAction ReqData to address error ")
		return
	}
	lastBlockHash, _ := base58.Decode(reqData.LastBlockHash)

	//获取私钥
	accounts, err := getSolPrivateKey(reqData.From)
	if err != nil {
		resp.Error(c, resp.FROM_ADDRESS)
		glog.Error("SolSignAction getSolPrivateKey error ", err)
		return
	}

	var instruction solana.Instruction
	if reqData.Token == "" {
		instruction = token.Transfer(fromAddress, toAddress, accounts[0].GetPubKey(), uint64(value))

	} else {
		tokenAddress, err := base58.Decode(reqData.Token)
		if err != nil {
			resp.Error(c, resp.PARAM_ERROR)
			glog.Error("SolSignAction ReqData token address error ")
			return
		}
		instruction = token.Transfer2(
			fromAddress,
			tokenAddress,
			toAddress,
			fromAddress, uint64(value), byte(reqData.Decimal))
	}
	tx := solana.NewTransaction(
		accounts[0].GetPubKey(),
		instruction,
	)

	tx.SetBlockhash(sha256.Sum256(lastBlockHash))
	err = tx.Sign(accounts[0].GetPrivKey())
	if err != nil {
		resp.Error(c, resp.SIGN_ERROR)
		glog.Error("SolSignAction Sign error ", err)
		return
	}

	signTx := base64.StdEncoding.EncodeToString(tx.Marshal())

	glog.Info("SolSignTx: ", signTx)
	resp.Ok(c, signTx)
}

func getSolPrivateKey(address string) ([]*wallet.SolWallet, error) {

	cModel, table := wallet.GetDBModel("sol")
	var addrM = &model.CoinAddress{}
	addrM.Address = address
	cModel.GetData(table, addrM)
	if addrM.PrivKey == "" {
		return nil, errors.New("No such address in the database.")
	}

	var accounts []*wallet.SolWallet

	s, _ := hex.DecodeString(addrM.PrivKey)
	w, err := wallet.NewSolWallet(s)
	if err != nil {
		return nil, err
	}
	accounts = append(accounts, w)

	return accounts, nil
}
