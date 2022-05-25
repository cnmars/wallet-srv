package wallet

import (
	"encoding/hex"
	crypto "wallet-srv/lib/pkg/xrp/crypto"
)

type XrpWallet struct {
	symbol     string
	privateKey []byte
	publicKey  []byte
}

func NewXrpWallet(seed []byte) (*XrpWallet, error) {

	privKey, err := crypto.NewECDSAKey(seed)
	if err != nil {
		return nil, err
	}

	sequenceZero := uint32(0)
	pubKey := privKey.Public(&sequenceZero)

	return &XrpWallet{
		symbol:     SymbolXrp,
		privateKey: seed,
		publicKey:  pubKey}, nil
}

func (w *XrpWallet) ChainId() int {
	return 0
}

func (w *XrpWallet) Symbol() string {
	return w.symbol
}

func (w *XrpWallet) DeriveAddress() string {

	key, err := crypto.NewECDSAKey(w.privateKey)
	if err != nil {
		return ""
	}
	sequenceZero := uint32(0)
	hash, err := crypto.AccountId(key, &sequenceZero)
	return hash.String()
}

func (w *XrpWallet) DerivePublicKey() string {
	return hex.EncodeToString(w.publicKey)
}

func (w *XrpWallet) DerivePrivateKey() string {
	return hex.EncodeToString(w.privateKey)
}

func (w *XrpWallet) DeriveNativePrivateKey() string {
	return ""
}
