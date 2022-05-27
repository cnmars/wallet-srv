package wallet

import (
	"github.com/golang/glog"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
)

type FlowWallet struct {
	symbol    string
	seed      []byte
	privKey   crypto.PrivateKey
	publicKey crypto.PublicKey
}

func NewFlowWallet(seed []byte) (*FlowWallet, error) {

	privKey, _ := crypto.GeneratePrivateKey(crypto.ECDSA_P256, seed)

	return &FlowWallet{
		symbol:    SymbolFlow,
		seed:      seed,
		privKey:   privKey,
		publicKey: privKey.PublicKey(),
	}, nil
}

func (w *FlowWallet) newAddress() (string, error) {

	address := flow.HexToAddress(w.DerivePublicKey())
	return "0x" + address.String(), nil
}

func (w *FlowWallet) ChainId() int {
	return 0
}

func (w *FlowWallet) Symbol() string {
	return w.symbol
}

func (w *FlowWallet) DeriveAddress() string {

	addr, err := w.newAddress()
	if err != nil {
		glog.Error(err)
		return ""
	}
	return addr
}

func (w *FlowWallet) DerivePublicKey() string {
	return w.publicKey.String()
}

func (w *FlowWallet) DerivePrivateKey() string {
	return w.privKey.String()
}

func (w *FlowWallet) DeriveNativePrivateKey() crypto.PrivateKey {
	return w.privKey
}

func (w *FlowWallet) GetPubKey() []byte {
	return nil
}

func (w *FlowWallet) GetPrivKey() []byte {
	return nil
}
