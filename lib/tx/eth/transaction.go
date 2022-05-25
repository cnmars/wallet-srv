package eth

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func SignTx(from common.Address, privateKey *ecdsa.PrivateKey, tx *types.Transaction, chainId *big.Int) (*types.Transaction, error) {

	signer := types.NewEIP155Signer(chainId)
	// Sign the transaction and verify the sender to avoid hardware fault surprises
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return nil, err
	}

	sender, err := types.Sender(signer, signedTx)
	if err != nil {
		return nil, err
	}

	if sender != from {
		return nil, fmt.Errorf("signer mismatch: expected %s, got %s", from.Hex(), sender.Hex())
	}

	return signedTx, nil
}
