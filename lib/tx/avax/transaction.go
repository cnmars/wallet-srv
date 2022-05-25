package avax

import (
	"fmt"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/codec/linearcodec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto"
	"github.com/ava-labs/avalanchego/utils/formatting"
	"github.com/ava-labs/avalanchego/utils/wrappers"
	"github.com/ava-labs/avalanchego/vms/avm"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

func setupCodec() (codec.GeneralCodec, codec.Manager) {
	c := linearcodec.NewDefault()
	m := codec.NewDefaultManager()
	errs := wrappers.Errs{}
	errs.Add(
		c.RegisterType(&avm.BaseTx{}),
		c.RegisterType(&secp256k1fx.TransferInput{}),
		c.RegisterType(&secp256k1fx.TransferOutput{}),
		c.RegisterType(&secp256k1fx.Credential{}),
		m.RegisterCodec(0, c),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
	return c, m
}

type Transaction struct {
	BaseTx avax.BaseTx
	Tx     *avm.Tx
}

func NewTransaction() *Transaction {
	return &Transaction{
		BaseTx: avax.BaseTx{
			NetworkID:    0,
			BlockchainID: ids.ID{},
			Ins:          []*avax.TransferableInput{},
			Outs:         []*avax.TransferableOutput{},
		},
		Tx: nil,
	}
}

func (btx *Transaction) AddInput(reqInputs []*Utxos) error {

	var ins []*avax.TransferableInput
	var tin *avax.TransferableInput
	for i, input := range reqInputs {
		txID, err := ids.FromString(input.TxId)
		if err != nil {
			continue
		}
		assetID, err := ids.FromString(input.AssetId)
		if err != nil {
			continue
		}
		tin = &avax.TransferableInput{
			UTXOID: avax.UTXOID{
				TxID:        txID,
				OutputIndex: uint32(input.Vout),
			},
			Asset: avax.Asset{ID: assetID},
			In: &secp256k1fx.TransferInput{
				Amt: input.Amt,
				Input: secp256k1fx.Input{
					SigIndices: []uint32{
						uint32(i),
					},
				},
			},
		}

		ins = append(ins, tin)
	}

	if len(ins) == 0 {
		return fmt.Errorf("no utxo")
	}
	btx.BaseTx.Ins = ins
	return nil
}

func (btx *Transaction) AddOutput(reqOutputs []*Outputs) error {

	var ous []*avax.TransferableOutput
	var outs *avax.TransferableOutput
	for _, output := range reqOutputs {

		assetId, err := ids.FromString(output.AssetId)
		if err != nil {
			continue
		}
		_, _, addr, err := formatting.ParseAddress(output.Address)
		if err != nil {
			continue
		}
		var shortIds []ids.ShortID
		addrShortId, _ := ids.ToShortID(addr)
		shortIds = append(shortIds, addrShortId)

		outs = &avax.TransferableOutput{
			Asset: avax.Asset{ID: assetId},
			Out: &secp256k1fx.TransferOutput{
				Amt: output.Amt,
				OutputOwners: secp256k1fx.OutputOwners{
					Threshold: 1,
					Addrs:     shortIds,
				},
			},
		}

		ous = append(ous, outs)
	}

	if len(ous) == 0 {
		return fmt.Errorf("no output")
	}

	btx.BaseTx.Outs = ous
	return nil
}

func (btx *Transaction) SetChainId(chainId uint32) {
	btx.BaseTx.NetworkID = chainId
}

func (btx *Transaction) SetBlockChainId(blockChainId string) {
	blockChainIds, err := ids.FromString(blockChainId)
	if err != nil {
		fmt.Println("blockChainId error: ", err)
		return
	}
	btx.BaseTx.BlockchainID = blockChainIds
}

//创建交易
func (btx *Transaction) Build() {
	btx.Tx = &avm.Tx{
		UnsignedTx: &avm.BaseTx{
			BaseTx: btx.BaseTx,
		},
	}
}

//交易签名
func (btx *Transaction) Sign(signers [][]*crypto.PrivateKeySECP256K1R) error {

	if btx.Tx == nil {
		return fmt.Errorf("btx.Tx not init.")
	}

	_, m := setupCodec()
	if err := btx.Tx.SignSECP256K1Fx(m, signers); err != nil {
		return err
	}

	return nil
}
