package action

import (
	"encoding/hex"
	"errors"
	"wallet-srv/lib"
	"wallet-srv/lib/pkg/ada/address"
	"wallet-srv/lib/pkg/ada/bip32"
	"wallet-srv/lib/pkg/ada/network"
	"wallet-srv/lib/resp"
	"wallet-srv/lib/tx/ada"
	"wallet-srv/lib/wallet"
	"wallet-srv/model"

	ctypes "wallet-srv/lib/types"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

//ADA验签
func AdaSignAction(c *gin.Context) {
	/*
		   {
			   "vin": [
				   {"TxHash": '', "Index": 0, "Amount": 500000},
				   {"TxHash": '', "Index": 0, "Amount": 500000},
				   {"TxHash": '', "Index": 0, "Amount": 500000},
			   ],
			   "vout": [
				   {"Address": 'xxxxx', "Amount": 100000},
				   {"Address": 'xxxxx', "Amount": 100000}
			   ],
			   "change": "xxxx",//找零地址
			   "fee_param": {
				   "txFeePerByte": 121,
				   "txFeeFixed": 1212,
				   "maxTxSize": 1212,
				   "protocolVersion": {
					   "major": 0,
					   "minor": 1,
				   },
				   "minUTxOValue": 1
			   }
		   }
		   **/
	var reqData ctypes.AdaTx
	err := lib.DecodeHTTPReqJSON(c, &reqData)
	if err != nil {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("AdaSignAction DecodeHTTPReqJSON error", err)
		return
	}

	if len(reqData.TxVin) == 0 {
		resp.Error(c, resp.REQ_ERROR)
		glog.Error("AdaSignAction DecodeHTTPReqJSON error", err)
		return
	}
	from := reqData.TxVin[0].Address
	privKey, err := getAdaPrivateKey(from)
	if err != nil {
		resp.Error(c, resp.FROM_ADDRESS)
		glog.Error("AdaSignAction getAdaPrivateKey error ", err)
		return
	}

	_, utxoPrvKey, _ := getBaseAddress(privKey)

	builder := ada.NewTxBuilder(
		reqData.FeeParam,
		[]bip32.XPrv{utxoPrvKey},
	)

	var sendAmount uint = 0
	for _, vout := range reqData.TxVout {
		sendAmount = ada.FloatToUint(vout.Amount)
	}

	txInputs := getTxInput(reqData.TxVin)
	var firstMatchInput ada.TxInput
	// Loop through utxos to find first input with enough ADA
	// TODO此处还可以优化 合并UTXO，归集UTXO等
	for _, utxo := range *txInputs {
		minRequired := sendAmount + 1000000 + 200000
		if utxo.Amount >= uint(minRequired) {
			firstMatchInput = utxo
		}
	}
	builder.AddInputs(&firstMatchInput)

	txOutputs := getTxOutput(reqData.TxVout)
	for _, vout := range *txOutputs {
		receiver, err := address.NewAddress(vout.Address.String())
		if err != nil {
			continue
		}
		builder.AddOutputs(ada.NewTxOutput(receiver, vout.Amount))
	}

	// Set TTL for 5 min into the future
	builder.SetTTL(uint32(reqData.Slot) + uint32(300))

	// Route back the change to the source address
	// This is equivalent to adding an output with the source address and change amount
	changeAddr, err := address.NewAddress(reqData.ChangeAddress)
	builder.AddChangeIfNeeded(changeAddr)

	// Build loops through the witness private keys and signs the transaction body hash
	signTx, err := builder.Build()
	if err != nil {
		resp.Error(c, resp.SIGN_ERROR)
		glog.Error("AdaSignAction sign error ", err)
		return
	}

	txid, _ := signTx.Hash()
	signTxHex, _ := signTx.Hex()
	var ret = &ctypes.Ret{
		TxId:   hex.EncodeToString(txid[:]),
		SignTx: signTxHex,
	}

	glog.Info("AdaSignTx: ", signTxHex)
	resp.Ok(c, ret)
}

//获取私钥
func getAdaPrivateKey(address string) (string, error) {

	cModel, table := wallet.GetDBModel("ada")
	var addrM = &model.CoinAddress{}
	addrM.Address = address
	cModel.GetData(table, addrM)

	if addrM.PrivKey == "" {
		return "", errors.New("No such address in the database.")
	}

	return addrM.PrivKey, nil
}

func getBaseAddress(seed string) (*address.BaseAddress, bip32.XPrv, error) {
	s, _ := hex.DecodeString(seed)
	rootKey := bip32.FromBip39Entropy(
		s,
		[]byte{},
	)

	accountKey := rootKey.Derive(address.Harden(1852)).Derive(address.Harden(1815)).Derive(address.Harden(0))

	utxoPrvKey := accountKey.Derive(0).Derive(0)
	utxoPubKey := utxoPrvKey.Public()
	utxoPubKeyHash := utxoPubKey.PublicKey().Hash()

	stakeKey := accountKey.Derive(2).Derive(0).Public()
	stakeKeyHash := stakeKey.PublicKey().Hash()

	baseAddr := address.NewBaseAddress(
		network.MainNet(),
		&address.StakeCredential{
			Kind:    address.KeyStakeCredentialType,
			Payload: utxoPubKeyHash[:],
		},
		&address.StakeCredential{
			Kind:    address.KeyStakeCredentialType,
			Payload: stakeKeyHash[:],
		})

	return baseAddr, utxoPrvKey, nil
}

func getTxInput(utxo []ctypes.Utxo) *[]ada.TxInput {
	var txinput []ada.TxInput
	for _, uxo := range utxo {
		txvin := &ada.TxInput{
			TxHash: []byte(uxo.Txid),
			Index:  uint16(uxo.Vout),
			Amount: ada.FloatToUint(uxo.Amount),
		}
		txinput = append(txinput, *txvin)
	}
	return &txinput
}

func getTxOutput(txvout []ctypes.Vout) *[]ada.TxOutput {

	var txoutput []ada.TxOutput
	for _, vout := range txvout {
		addr, err := address.NewAddress(vout.Address)
		if err != nil {
			continue
		}
		txout := &ada.TxOutput{
			Address: addr,
			Amount:  ada.FloatToUint(vout.Amount),
		}
		txoutput = append(txoutput, *txout)
	}

	return &txoutput
}
