package xmr

import (
	"wallet-srv/lib/pkg/monero"
	"wallet-srv/lib/types"
)

type Tx struct {
	*monero.Transaction
}

func NewTx() *Tx {
	t := new(monero.Transaction)
	return &Tx{
		t,
	}
}

func (tx *Tx) AddInput(vin []*types.Utxo) {
	var txInkey []*monero.TxInToKey
	for _, in := range vin {
		txKey := &monero.TxInToKey{
			amount: in.Amount,
		}
	}

}
