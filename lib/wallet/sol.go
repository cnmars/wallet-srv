package wallet

import (
	"crypto/ed25519"
	"encoding/hex"

	"github.com/btcsuite/btcutil/base58"
)

type SolWallet struct {
	symbol    string
	privKey   ed25519.PrivateKey
	publicKey ed25519.PublicKey
}

func NewSolWallet(seed []byte) (*SolWallet, error) {

	privKey := ed25519.NewKeyFromSeed(seed)

	publicKey := make([]byte, ed25519.PublicKeySize)
	copy(publicKey, privKey[32:])

	return &SolWallet{
		symbol:    SymbolSol,
		privKey:   privKey,
		publicKey: publicKey,
	}, nil
}

func (w *SolWallet) ChainId() int {
	return 0
}

func (w *SolWallet) Symbol() string {
	return w.symbol
}

func (w *SolWallet) DeriveAddress() string {
	return w.toBase58()
}

func (w *SolWallet) DerivePublicKey() string {
	return hex.EncodeToString(w.publicKey)
}

func (w *SolWallet) toBase58() string {
	return base58.Encode(w.publicKey)
}

func (w *SolWallet) DerivePrivateKey() string {
	return hex.EncodeToString(w.privKey)
}

func (w *SolWallet) DeriveNativePrivateKey() *ed25519.PrivateKey {
	return &w.privKey
}

func (w *SolWallet) GetPubKey() ed25519.PublicKey {
	return w.publicKey
}

func (w *SolWallet) GetPrivKey() ed25519.PrivateKey {
	return w.privKey
}
