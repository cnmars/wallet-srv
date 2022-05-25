package wallet

import (
	"encoding/hex"

	"github.com/ava-labs/avalanchego/utils/crypto"
	"github.com/ava-labs/avalanchego/utils/formatting"
	"github.com/golang/glog"
)

type AvaxWallet struct {
	symbol    string
	seed      []byte
	privKey   crypto.PrivateKey
	publicKey crypto.PublicKey
}

func NewAvaxWallet(seed []byte) (*AvaxWallet, error) {
	factory := crypto.FactorySECP256K1R{}
	privKey, _ := factory.ToPrivateKey(seed)
	return &AvaxWallet{
		symbol:    SymbolAvax,
		seed:      seed,
		privKey:   privKey,
		publicKey: privKey.PublicKey(),
	}, nil
}

func (w *AvaxWallet) newAddress() (string, error) {
	address, _ := formatting.FormatAddress("x", "avax", w.publicKey.Address().Bytes())
	return address, nil
}

func (w *AvaxWallet) ChainId() int {
	return 0
}

func (w *AvaxWallet) Symbol() string {
	return w.symbol
}

func (w *AvaxWallet) DeriveAddress() string {

	addr, err := w.newAddress()
	if err != nil {
		glog.Error(err)
		return ""
	}
	return addr
}

func (w *AvaxWallet) DerivePublicKey() string {
	return hex.EncodeToString(w.publicKey.Bytes())
}

func (w *AvaxWallet) DerivePrivateKey() string {
	return hex.EncodeToString(w.privKey.Bytes())
}

func (w *AvaxWallet) DeriveNativePrivateKey() crypto.PrivateKey {
	return w.privKey
}

func (w *AvaxWallet) GetPubKey() []byte {
	return w.publicKey.Bytes()
}

func (w *AvaxWallet) GetPrivKey() []byte {
	return w.privKey.Bytes()
}
