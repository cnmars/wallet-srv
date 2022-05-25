package avax

type Utxos struct {
	AssetId string `json:"asset_id"`
	TxId    string `json:"txid"`
	Vout    uint8  `json:"vout"`
	Amt     uint64 `json:"amount"`
	Address string `json:"address"`
}

type Outputs struct {
	AssetId string `json:"asset_id"`
	Address string `json:"address"`
	Amt     uint64 `json:"amount"`
}
