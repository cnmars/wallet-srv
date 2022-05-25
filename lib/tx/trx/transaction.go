package trx

import (
	"crypto/ecdsa"
	"crypto/sha256"

	"github.com/ethereum/go-ethereum/crypto"
)

//交易签名
func SignTx(privateKey *ecdsa.PrivateKey, rawData []byte) ([]byte, error) {

	txHash := sha256.Sum256(rawData)

	signTx, err := crypto.Sign(txHash[:], privateKey)
	if err != nil {
		return nil, err
	}

	return signTx, nil
}
