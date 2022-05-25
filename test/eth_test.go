package eth_test

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

	hdw, err := wallet.NewHDWallet(wallet.SymbolEth, wallet.ETH)
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

func TestEthSign(t *testing.T) {

	from := "0x7eeE959B97a243233EDd1133c7D706Be97999D49"
	to := "0x09d247829344c4d1D1A623E01C9387fF978E1cEe"
	amount := int64(1000)
	nonce := uint64(0)
	gasLimit := uint64(21000)
	gasPrice := int64(18 * wallet.WeiPerGwei)
	chainID := int64(1)
	contract := ""

	ethTx := &types.EthTx{
		From:     from,
		To:       to,
		Amount:   amount,
		Nonce:    nonce,
		GasLimit: gasLimit,
		GasPrice: gasPrice,
		ChainID:  chainID,
		Contract: contract,
	}

	client := rpctest.NewClient()
	var signHex rpctest.RespSignHex
	err := rpctest.DoSign(client, "/eth/sign", ethTx, &signHex)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	t.Errorf("tx -> %v, err:%v", signHex.Content, err)

}
