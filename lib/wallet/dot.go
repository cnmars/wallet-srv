package wallet

import (
	"encoding/hex"
	"fmt"
	"wallet-srv/lib/pkg/dot/signature"
)

type DotWallet struct {
	symbol     string
	privateKey []byte
	publicKey  []byte
}

const DOT_NETWORK = 0

func NewDotWallet(seed []byte) (*DotWallet, error) {

	key, err := signature.KeyringPairFromSecret(getUri(seed), DOT_NETWORK)
	if err != nil {
		fmt.Println("NewDotWallet: ", err)
		return nil, err
	}

	return &DotWallet{
		symbol:     SymbolDot,
		privateKey: seed,
		publicKey:  key.PublicKey}, nil
}

func (w *DotWallet) ChainId() int {
	return 0
}

func (w *DotWallet) Symbol() string {
	return w.symbol
}

func (w *DotWallet) DeriveAddress() string {

	key, err := signature.KeyringPairFromSecret(getUri(w.privateKey), DOT_NETWORK)
	if err != nil {
		return ""
	}

	return key.Address
}

func (w *DotWallet) DerivePublicKey() string {
	return hex.EncodeToString(w.publicKey)
}

func (w *DotWallet) DerivePrivateKey() string {
	return hex.EncodeToString(w.privateKey)
}

func (w *DotWallet) DeriveNativePrivateKey() string {
	return ""
}

func (w *DotWallet) GetKeys() (*signature.KeyringPair, error) {
	key, err := signature.KeyringPairFromSecret(getUri(w.privateKey), DOT_NETWORK)
	if err != nil {
		fmt.Println("NewDotWallet: ", err)
		return nil, err
	}
	return &key, nil
}

func getUri(seed []byte) string {
	return "//" + hex.EncodeToString(seed)
}
