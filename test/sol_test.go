package sol_test

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

func TestAddress(t *testing.T) {

	hdw, err := wallet.NewHDWallet(wallet.SymbolSol, wallet.SOL)
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

func TestSolSign(t *testing.T) {

	from := "GwYeRRgRhykYRX9taXWVALXbCJQaUECMG4s6gHS4sxUv"
	to := "EWmowLYfz49uyNJdT22usS9qQm495qZdKSYsGQjLw6PJ"
	amount := int64(10000)
	lastBlockHash := "6MEVqZdfpkunirTp1UHXecYjJmfxuFLwhTQKai7LePKd"

	solTx := &types.SolTx{
		From:          from,
		To:            to,
		Amount:        amount,
		LastBlockHash: lastBlockHash,
	}

	client := rpctest.NewClient()
	var signHex rpctest.RespSignHex
	err := rpctest.DoSign(client, "/sol/sign", solTx, &signHex)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	t.Errorf("tx -> %v, err:%v", signHex.Content, err)
}
