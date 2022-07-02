package wallet

import (
	"encoding/hex"
	"wallet-srv/lib/pkg/monero"
)

type XmrWallet struct {
	symbol     string
	seed       []byte
	privateKey *monero.Key
	publicKey  *monero.Key
}

func NewXmrWallet(seed []byte) (*XmrWallet, error) {

	key := monero.NewKey(seed)
	return &XmrWallet{
		symbol:     SymbolXmr,
		seed:       seed,
		privateKey: key,
		publicKey:  key.PubKey()}, nil
}

func (w *XmrWallet) ChainId() int {
	return 0
}

func (w *XmrWallet) Symbol() string {
	return w.symbol
}

func (w *XmrWallet) DeriveAddress() string {
	key := monero.NewKey(w.seed)
	return key.Address()
}

func (w *XmrWallet) DerivePublicKey() string {
	return hex.EncodeToString(w.publicKey.Serialize())
}

func (w *XmrWallet) DerivePrivateKey() string {
	return hex.EncodeToString(w.privateKey.Serialize())
}

func (w *XmrWallet) DeriveNativePrivateKey() string {
	return ""
}
