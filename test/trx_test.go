package trx_test

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

	hdw, err := wallet.NewHDWallet(wallet.SymbolTrx, wallet.TRX)
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

func TestTrxSign(t *testing.T) {
	/*
		{"visible":true,"txID":"a16628a45e4d600bd51fd831b0c64cc7558c4d610ca94ba12cce08d4a53fb1ed","raw_data":{"contract":[{"parameter":{"value":{"amount":10000,"owner_address":"THWUyyjjWY5oAc4KBK3QZKZgGJtnS41d2r","to_address":"TWktyDFULE3vsEnR4gDwRrxd2PVhRVUhQY"},"type_url":"type.googleapis.com\/protocol.TransferContract"},"type":"TransferContract"}],"ref_block_bytes":"cf4e","ref_block_hash":"fcad8a445616ec30","expiration":1648478124000,"timestamp":1648478066193},"raw_data_hex":"0a02cf4e2208fcad8a445616ec3040e0efcf87fd2f5a66080112620a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412310a154152b31daab5a836bef19c5f2c6aad8ff76dc76180121541e4069902588f1111e28c0f90ac696700be7cd57918904e7091accc87fd2f"}
	*/
	from := "THWUyyjjWY5oAc4KBK3QZKZgGJtnS41d2r"
	rawDataHex := "0a02cf4e2208fcad8a445616ec3040e0efcf87fd2f5a66080112620a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412310a154152b31daab5a836bef19c5f2c6aad8ff76dc76180121541e4069902588f1111e28c0f90ac696700be7cd57918904e7091accc87fd2f"

	trxTx := &types.TrxTx{
		From:       from,
		RawDataHex: rawDataHex,
	}

	client := rpctest.NewClient()
	var signHex rpctest.RespSignHex
	err := rpctest.DoSign(client, "/trx/sign", trxTx, &signHex)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	t.Errorf("tx -> %v, err:%v", signHex.Content, err)
}
