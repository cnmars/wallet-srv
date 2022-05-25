package btc_test

import (
	"flag"
	"testing"
	"wallet-srv/lib/types"
	"wallet-srv/lib/wallet"
	"wallet-srv/test/rpctest"
)

func init() {
	flag.Lookup("alsologtostderr").Value.Set("true")
}
func TestBtcSign(t *testing.T) {
	a := "mrXUCNgmV439zQjh13N4kSG6DJuqJ9zD4C"
	a1 := "mvaCNagSKWA2cpyXoejJCEAFRdnaPeQ3SD"
	feePerKb := int64(80 * 1000)
	changeAddress := "mrXUCNgmV439zQjh13N4kSG6DJuqJ9zD4C"

	vin := types.Utxo{Address: a, Txid: "c1076d23d1d92170c01588eace8dfedf5284888402fe79344b796749d5eb19a1", Amount: 0.001, Vout: 0}
	txvin := make([]types.Utxo, 0)
	txvin = append(txvin, vin)

	txvout := make([]types.Vout, 0)
	vout := types.Vout{Address: a1, Amount: 0.0002}
	txvout = append(txvout, vout)
	change := types.Vout{Address: a, Amount: 0.0005}
	txvout = append(txvout, change)

	btcTx := &types.BtcTx{
		TxVin:         txvin,
		TxVout:        txvout,
		FeePerKb:      feePerKb,
		ChangeAddress: changeAddress,
	}

	client := rpctest.NewClient()
	var signHex rpctest.RespSignHex
	err := rpctest.DoSign(client, "/btc/sign", btcTx, &signHex)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	t.Errorf("tx -> %v, err:%v", signHex.Content, err)
}

func TestAddress(t *testing.T) {

	hdw, err := wallet.NewHDWallet(wallet.SymbolBtc, wallet.BTC)
	if err != nil {
		t.Errorf("TestAddress NewHDWallet error:%v", err)
		return
	}

	w, err := hdw.NewWallet(0, 0, 0)
	if err != nil {
		t.Errorf("TestAddress NewWallet error:%v", err)
		return
	}

	t.Errorf("\nseed->%v \naddress->%v \npublicKey->%v \nprivkey->%v", hdw.GetSeed(), w.DeriveAddress(), w.DerivePublicKey(), w.DerivePrivateKey())
}
