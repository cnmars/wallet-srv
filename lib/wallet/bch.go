package wallet

import (
	"wallet-srv/lib/bchutil"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func GetBitCoinCashAddress(address string) (string, error) {
	chainParams := &chaincfg.MainNetParams
	addr, err := btcutil.DecodeAddress(address, chainParams)
	if err != nil {
		return "", err
	}

	addrHash, err := bchutil.NewCashAddressPubKeyHash(addr.ScriptAddress(), chainParams)
	if err != nil {
		return "", err
	}

	return addrHash.EncodeAddress(), nil
}
