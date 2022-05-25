package luna

import (
	"fmt"
	"wallet-srv/lib/pkg/terra/msg"
	"wallet-srv/lib/pkg/terra/params"
	"wallet-srv/lib/pkg/terra/tx"

	cctypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

// TxOptions tx creation options
type TxOptions struct {
	Msgs []msg.Msg
	Memo string

	// Optional parameters
	AccountNumber uint64
	Sequence      uint64
	GasLimit      uint64
	FeeAmount     msg.Coins

	SignMode      tx.SignMode
	FeeGranter    msg.AccAddress
	TimeoutHeight uint64
}

var EncodingConfig = params.MakeEncodingConfig()

// CreateAndSignTx build and sign tx
func NewTxSign(options TxOptions, privKey cctypes.PrivKey) (*tx.Builder, error) {

	txbuilder := tx.NewTxBuilder(EncodingConfig.TxConfig)
	txbuilder.SetFeeAmount(options.FeeAmount)
	txbuilder.SetFeeGranter(options.FeeGranter)
	txbuilder.SetGasLimit(options.GasLimit)
	txbuilder.SetMemo(options.Memo)
	txbuilder.SetMsgs(options.Msgs...)
	txbuilder.SetTimeoutHeight(options.TimeoutHeight)

	// use direct sign mode as default
	if tx.SignModeUnspecified == options.SignMode {
		options.SignMode = tx.SignModeDirect
	}

	if options.AccountNumber == 0 || options.Sequence == 0 {
		return nil, fmt.Errorf("from account is not empty.")
	}
	gasLimit := int64(options.GasLimit)
	txbuilder.SetGasLimit(uint64(gasLimit))
	if options.GasLimit == 0 {
		return nil, fmt.Errorf("gasLimit not empty.")
	}

	if options.FeeAmount.IsZero() {
		return nil, fmt.Errorf("FeeAmout is not 0.")
	}

	err := txbuilder.Sign(options.SignMode, tx.SignerData{
		AccountNumber: options.AccountNumber,
		ChainID:       params.ChainID,
		Sequence:      options.Sequence,
	}, privKey, true)
	if err != nil {
		return nil, fmt.Errorf("failed to sign tx")
	}

	return &txbuilder, nil
}
