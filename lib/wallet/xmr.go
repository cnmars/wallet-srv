package wallet

import (
	"encoding/hex"
	"wallet-srv/lib/pkg/monero"
)

type XmrWallet struct {
	symbol     string
	privateKey []byte
	publicKey  []byte
}

func NewXmrWallet(seed []byte) (*XmrWallet, error) {

	key := monero.NewKey(seed)

	return &XmrWallet{
		symbol:     SymbolXmr,
		privateKey: seed,
		publicKey:  key.PubKey()[:]}, nil
}

func (w *XmrWallet) ChainId() int {
	return 0
}

func (w *XmrWallet) Symbol() string {
	return w.symbol
}

func (w *XmrWallet) DeriveAddress() string {
	key := monero.NewKey(w.privateKey)
	return key.Address()
}

func (w *XmrWallet) DerivePublicKey() string {
	return hex.EncodeToString(w.publicKey)
}

func (w *XmrWallet) DerivePrivateKey() string {
	key := monero.NewKey(w.privateKey)
	return hex.EncodeToString(key[:])
}

func (w *XmrWallet) DeriveNativePrivateKey() string {
	return ""
}
