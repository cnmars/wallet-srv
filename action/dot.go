package action

import (
	"encoding/hex"
	"errors"
	"fmt"
	"wallet-srv/lib"
	"wallet-srv/lib/pkg/dot"
	"wallet-srv/lib/pkg/dot/signature"
	"wallet-srv/lib/pkg/dot/types"
	"wallet-srv/lib/resp"
	ctypes "wallet-srv/lib/types"
	"wallet-srv/lib/wallet"
	"wallet-srv/model"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func DotSignAction(c *gin.Context) {
	/*
		   {
			   "from": "",
			   "to": "",
			   "amount": 12310201,
			   "fee": 0,
			   "last_blockhash": "xxxxxx",
			   "nonce": 12,
			   "spec_ver": 9090,
			   "tran_ver":14
		   }
		   **/
	var reqData ctypes.DotTx
	err := lib.DecodeHTTPReqJSON(c, &reqData)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("DotSignAction: DecodeHTTPReqJSON error", err)
		return
	}

	w, err := getDotPrivateKey(reqData.From)
	if err != nil {
		resp.Error(c, resp.PARAM_ERROR)
		glog.Error("DotSignAction: getData error ", err)
		return
	}

	toAddr := types.NewAddressFromAccountID(addressToAccountId(reqData.To))
	amount := types.NewUCompactFromUInt(reqData.Amount)
	fee := types.NewUCompactFromUInt(reqData.Fee)

	//cache local meta
	//exec command get meta cache file ./wallet-tool -dot-meta -c ../config/config.json
	meta, err := dot.GetMeta()
	if err != nil {
		resp.Error(c, resp.SERVER_INNER_ERROR)
		glog.Error("DotSignAction: dot.GetMeta() ", err)
		return
	}

	call, err := types.NewCall(meta, "Balances.transfer", toAddr, amount)
	if err != nil {
		resp.Error(c, resp.SERVER_INNER_ERROR)
		glog.Error("DotSignAction: types.NewCall ", err)
		return
	}

	// Create the extrinsic
	ext := types.NewExtrinsic(call)

	nonce := uint32(reqData.Nonce)
	genesisHash := types.NewHash(types.MustHexDecodeString(reqData.LastBlockHash))
	o := types.SignatureOptions{
		BlockHash:          genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(nonce)),
		SpecVersion:        types.NewU32(reqData.SpecVer),
		Tip:                fee,
		TransactionVersion: types.NewU32(reqData.TranVersion),
	}

	// Sign the transaction using Alice's default account
	key, err := w.GetKeys()
	if err != nil {
		resp.Error(c, resp.SERVER_INNER_ERROR)
		glog.Error("DotSignAction:  w.GetKeys ", err)
		return
	}
	err = ext.Sign(*key, o)
	if err != nil {
		resp.Error(c, resp.SIGN_ERROR)
		glog.Error("DotSignAction:  ext.Sign ", err)
		return
	}

	tx, _ := ext.MarshalJSON()
	txHex := hex.EncodeToString(tx)
	glog.Info("DotSignTxHex: ", txHex)

	resp.Ok(c, txHex)
}

func getDotPrivateKey(address string) (*wallet.DotWallet, error) {
	cModel, table := wallet.GetDBModel("dot")
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
	w, err := wallet.NewDotWallet(s)
	if err != nil {
		return nil, fmt.Errorf("NewDotWallet fail by privKey")
	}
	return w, nil
}

func addressToAccountId(addr string) []byte {
	return signature.SS58AddressToPublicKey(addr)
}
