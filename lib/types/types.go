package types

import (
	"wallet-srv/lib/pkg/ada/protocol"
	"wallet-srv/lib/tx/avax"
)

type Ret struct {
	TxId   string
	SignTx string
}

type Vout struct {
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
}

type Utxo struct {
	Address string  `json:"address"`
	Txid    string  `json:"txid"`
	Vout    uint32  `json:"vout"`
	Amount  float64 `json:"amount"`
}

type BtcTx struct {
	TxVin         []Utxo `json:"vin"`
	TxVout        []Vout `json:"vout"`
	ChangeAddress string `json:"change"`
	FeePerKb      int64  `json:"feekb"`
}

type HdTx struct {
	TxVin         []Utxo `json:"vin"`
	TxVout        []Vout `json:"vout"`
	ChangeAddress string `json:"change"`
	FeePerKb      int64  `json:"feekb"`
	Coin          string `json:"coin"`
}

type EthTx struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Amount   int64  `json:"amount"`
	Nonce    uint64 `json:"nonce"`
	GasLimit uint64 `json:"gaslimit"`
	GasPrice int64  `json:"gasprice"`
	Contract string `json:"contract"`
	ChainID  int64  `json:"chainid"`
}

type TrxTx struct {
	From       string `json:"from"`
	RawDataHex string `json:"raw_data_hex"`
}

type SolTx struct {
	From          string `json:"from"`
	To            string `json:"to"`
	Amount        int64  `json:"amount"`
	LastBlockHash string `json:"last_blockhash"`
	Token         string `json:"token"`
	Decimal       uint8  `json:"decimal"`
}

type FilTx struct {
	From       string  `json:"from"`
	To         string  `json:"to"`
	Amount     float64 `json:"amount"`
	Nonce      uint64  `json:"nonce"`
	GasLimit   uint64  `json:"gaslimit"`
	GasFee     int64   `json:"gasfee"`
	GasPremium int64   `json:"gaspremium"`
	Param      string  `json:"param"`
}

type AdaTx struct {
	TxVin         []Utxo            `json:"vin"`
	TxVout        []Vout            `json:"vout"`
	ChangeAddress string            `json:"change"`
	FeeParam      protocol.Protocol `json:"fee_param"`
	Slot          uint              `json:"last_slot"`
}

type XrpTx struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Amount      string `json:"amount"`
	AssetCode   string `json:"assetcode"`
	AssetIssuer string `json:"assetissuer"`
	Fee         string `json:"fee"`
	FromSeq     uint32 `json:"sequence"`
}

type LunaTx struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Amount    string `json:"amount"`
	FeeAmount string `json:"fee"`
	GasLimit  uint64 `json:"gaslimit"`
	Seq       uint64 `json:"sequence"`
	Coin      string `json:"coin"`
}

type XmrTx struct {
	TxHash string `json:"tx_hash"`
	TxKey  string `json:"tx_key"`
	From   string `json:"from"`
	TxHex  string `json:"tx_hex"`
}

type NearTx struct {
	From          string  `json:"from"`
	To            string  `json:"to"`
	Amount        float64 `json:"amount"`
	LastBlockHash string  `json:"last_blockhash"`
	Nonce         int64   `json:"nonce"`
}

type DotTx struct {
	From          string `json:"from"`
	To            string `json:"to"`
	Amount        uint64 `json:"amount"`
	Fee           uint64 `json:"fee"`
	LastBlockHash string `json:"last_blockhash"`
	Nonce         int64  `json:"nonce"`
	SpecVer       uint32 `json:"spec_ver"`
	TranVersion   uint32 `json:"tran_ver"`
}

type AvaxTx struct {
	ChainId      uint32          `json:"chain_id"`
	BlockChainId string          `json:"block_chainid"`
	Inputs       []*avax.Utxos   `json:"vin"`
	Outputs      []*avax.Outputs `json:"vout"`
	Memo         string          `json:"memo"`
}

type ReqAddress struct {
	Coin  string `json:"coin"`
	Limit uint32 `json:"limit"`
}
