package fil

import (
	"fmt"
	"wallet-srv/lib/filutil"
	"wallet-srv/lib/pkg/secp256k1"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/golang/glog"
	"github.com/minio/blake2b-simd"
	"github.com/shopspring/decimal"
)

// CreateTransaction func
func CreateTransaction(from, to filutil.Address, val float64, gaslimit int64, gasfee, gaspremium int64, nonce uint64,
	method uint64, params []byte) (tx *filutil.Message, err error) {

	tx = &filutil.Message{
		Version:    0,
		To:         to,
		From:       from,
		Nonce:      0,
		Value:      filutil.FromFil(decimal.NewFromFloat(val)),
		GasLimit:   gaslimit,
		GasFeeCap:  abi.NewTokenAmount(gasfee),
		GasPremium: abi.NewTokenAmount(gaspremium),
		Method:     method,
		Params:     params,
	}

	return tx, nil
}

// WalletSignMessage signs the given message using the given private key.
func SignMessage(pk []byte, msg *filutil.Message) (*filutil.SignedMessage, error) {
	mb, err := msg.ToStorageBlock()
	if err != nil {
		return nil, fmt.Errorf("serializing message: %w", err)
	}

	sig, err := sign(pk, mb.Cid().Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to sign message: %w", err)
	}

	return &filutil.SignedMessage{
		Message:   msg,
		Signature: sig,
	}, nil
}

func VerifyMessage(sm *filutil.SignedMessage) bool {
	err := verify(sm.Signature.Data, sm.Message.From, sm.Message.Cid().Bytes())

	if err != nil {
		glog.Errorf("VerifyMessage error: ", err)
		return false
	}
	return true
}

func sign(privKey, msg []byte) (*crypto.Signature, error) {
	b2sum := blake2b.Sum256(msg)
	sig, err := secp256k1.Sign(privKey, b2sum[:])
	if err != nil {
		return nil, err
	}
	return &crypto.Signature{
		Type: crypto.SigTypeSecp256k1,
		Data: sig,
	}, nil
}

func verify(sig []byte, a filutil.Address, msg []byte) error {
	b2sum := blake2b.Sum256(msg)
	pubk, err := secp256k1.EcRecover(b2sum[:], sig)
	if err != nil {
		return err
	}

	maybeaddr, err := filutil.NewSecp256k1Address(pubk)
	if err != nil {
		return err
	}

	if a != maybeaddr {
		return fmt.Errorf("signature did not match")
	}

	return nil
}
