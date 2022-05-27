package action

import (
	"encoding/hex"
	"errors"

	"wallet-srv/lib"

	"wallet-srv/lib/resp"
	"wallet-srv/lib/wallet"

	cflow "wallet-srv/lib/tx/flow"
	ctypes "wallet-srv/lib/types"
	"wallet-srv/model"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/onflow/flow-go-sdk"
)

func FlowSignAction(c *gin.Context) {
	/*
		{
			"from": "",
			"to": "",
			"player": "",
			"amount": "102.99911",
			"last_blockhash": "" //last block hash
			"token": "1654653399040a61",
			"seq_num": 121
		}
		**/
	var reqData ctypes.FlowTx
	err := lib.DecodeHTTPReqJSON(c, &reqData)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("SolSignAction: DecodeHTTPReqJSON error", err)
		return
	}

	value := reqData.Amount
	sender := flow.HexToAddress(reqData.From)
	player := flow.HexToAddress(reqData.Player)

	//获取私钥
	senderW, err := getFlowPrivateKey(reqData.From)
	if err != nil {
		resp.Error(c, resp.FROM_ADDRESS)
		glog.Error("FlowSignAction from privkey error ", err)
		return
	}
	playerW, err := getFlowPrivateKey(reqData.Player)
	if err != nil {
		resp.Error(c, resp.FROM_ADDRESS)
		glog.Error("FlowSignAction player privkey error ", err)
		return
	}

	tx := cflow.NewTx()
	tx.SetSender(sender)
	tx.SetPlayer(player)
	tx.SetAuthorizer(player)
	tx.SetAuthorizer(sender)
	tx.SetProposalKey(player, 0, reqData.FromSeqNum)
	tx.SetFeeLimit(100)
	tx.SetTransferScript()
	tx.SetParam(value, reqData.To)
	tx.SetBlockHash(reqData.LastBlockHash)

	err = tx.Sign(senderW.DeriveNativePrivateKey(), playerW.DeriveNativePrivateKey())
	if err != nil {
		resp.Error(c, resp.SIGN_ERROR)
		glog.Error("FlowSignAction sign error ", err)
		return
	}

	signHexTx := hex.EncodeToString(tx.Tx.Encode())

	glog.Info("FlowSignTx: ", signHexTx)
	resp.Ok(c, signHexTx)
}

func getFlowPrivateKey(address string) (*wallet.FlowWallet, error) {

	cModel, table := wallet.GetDBModel("flow")
	var addrM = &model.CoinAddress{}
	addrM.Address = address
	cModel.GetData(table, addrM)
	if addrM.PrivKey == "" {
		return nil, errors.New("No such address in the database.")
	}

	w, err := wallet.NewFlowWallet([]byte(addrM.PrivKey))
	if err != nil {
		return nil, err
	}

	return w, nil
}
