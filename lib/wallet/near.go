package wallet

import (
	"encoding/hex"
	"wallet-srv/lib/pkg/near/account"
)

type NearWallet struct {
	symbol     string
	privateKey []byte
	publicKey  []byte
}

func NewNearWallet(seed []byte) (*NearWallet, error) {

	_, pubKey, _ := account.GenerateKeys(seed)

	return &NearWallet{
		symbol:     SymbolNear,
		privateKey: seed,
		publicKey:  pubKey}, nil
}

func (w *NearWallet) ChainId() int {
	return 0
}

func (w *NearWallet) Symbol() string {
	return w.symbol
}

func (w *NearWallet) DeriveAddress() string {
	return account.PublicKeyToAddress(w.publicKey)
}

func (w *NearWallet) DerivePublicKey() string {
	return hex.EncodeToString(w.publicKey)
}

func (w *NearWallet) DerivePrivateKey() string {
	return hex.EncodeToString(w.privateKey)
}

func (w *NearWallet) DeriveNativePrivateKey() string {
	return ""
}
