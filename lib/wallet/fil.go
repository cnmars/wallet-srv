package wallet

import (
	"encoding/hex"
	"wallet-srv/lib/filutil"
	"wallet-srv/lib/pkg/secp256k1"

	"github.com/golang/glog"
)

type FilWallet struct {
	symbol    string
	privKey   []byte
	publicKey []byte
}

func NewFilWallet(seed []byte) (*FilWallet, error) {

	pk := secp256k1.PublicKey(seed)

	return &FilWallet{
		symbol:    SymbolSol,
		privKey:   seed,
		publicKey: pk,
	}, nil
}

func (w *FilWallet) newAddress() (string, error) {

	//使用公钥地址
	addr, err := filutil.NewSecp256k1Address(w.publicKey)
	if err != nil {
		return "", err
	}

	return addr.String(), nil
}

func (w *FilWallet) ChainId() int {
	return 0
}

func (w *FilWallet) Symbol() string {
	return w.symbol
}

func (w *FilWallet) DeriveAddress() string {

	addr, err := w.newAddress()
	if err != nil {
		glog.Error(err)
		return ""
	}
	return addr
}

func (w *FilWallet) DerivePublicKey() string {
	return hex.EncodeToString(w.publicKey)
}

func (w *FilWallet) DerivePrivateKey() string {
	return hex.EncodeToString(w.privKey)
}

func (w *FilWallet) DeriveNativePrivateKey() []byte {
	return w.privKey
}

func (w *FilWallet) GetPubKey() []byte {
	return w.publicKey
}

func (w *FilWallet) GetPrivKey() []byte {
	return w.privKey
}
