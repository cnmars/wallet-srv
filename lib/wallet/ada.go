package wallet

import (
	"wallet-srv/lib/pkg/ada/address"
	"wallet-srv/lib/pkg/ada/bip32"
	"wallet-srv/lib/pkg/ada/network"
)

type AdaWallet struct {
	symbol    string
	privKey   []byte
	publicKey []byte
}

func NewAdaWallet(seed []byte) (*AdaWallet, error) {

	rootKey := bip32.FromBip39Entropy(
		seed,
		[]byte{},
	)

	return &AdaWallet{
		symbol:    SymbolSol,
		privKey:   seed,
		publicKey: rootKey,
	}, nil
}

func (w *AdaWallet) ChainId() int {
	return 0
}

func (w *AdaWallet) Symbol() string {
	return w.symbol
}

func (w *AdaWallet) DeriveAddress() string {
	rootKey := bip32.FromBip39Entropy(
		w.privKey,
		[]byte{},
	)

	accountKey := rootKey.Derive(address.Harden(1852)).Derive(address.Harden(1815)).Derive(address.Harden(0))

	utxoPubKey := accountKey.Derive(0).Derive(0).Public()
	utxoPubKeyHash := utxoPubKey.PublicKey().Hash()

	stakeKey := accountKey.Derive(2).Derive(0).Public()
	stakeKeyHash := stakeKey.PublicKey().Hash()

	baseAddr := address.NewBaseAddress(
		network.MainNet(),
		&address.StakeCredential{
			Kind:    address.KeyStakeCredentialType,
			Payload: utxoPubKeyHash[:],
		},
		&address.StakeCredential{
			Kind:    address.KeyStakeCredentialType,
			Payload: stakeKeyHash[:],
		})
	return baseAddr.ToEnterprise().String()
}

func (w *AdaWallet) DerivePublicKey() string {
	return ""
}

func (w *AdaWallet) DerivePrivateKey() string {
	return ""
}

func (w *AdaWallet) DeriveNativePrivateKey() []byte {
	return w.privKey
}

func (w *AdaWallet) GetPubKey() []byte {
	return w.publicKey
}

func (w *AdaWallet) GetPrivKey() []byte {
	return w.privKey
}
