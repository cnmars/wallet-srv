package wallet

import (
	"github.com/btcsuite/btcd/chaincfg"
)

// wallet type from bip44
const (
	Zero      uint32 = 0
	ZeroQuote uint32 = 0

	// https://github.com/satoshilabs/slips/blob/master/slip-0044.md#registered-coin-types
	BTC        = ZeroQuote + 0
	BTCTestnet = ZeroQuote + 1
	LTC        = ZeroQuote + 2
	DOGE       = ZeroQuote + 3
	DASH       = ZeroQuote + 5
	TRX        = ZeroQuote + 41
	ETH        = ZeroQuote + 60
	BCH        = ZeroQuote + 145
	QTUM       = ZeroQuote + 2301
	SOL        = ZeroQuote + 501
	FIL        = ZeroQuote + 461
	LUNA       = ZeroQuote + 330
	ADA        = ZeroQuote + 1815
	XRP        = ZeroQuote + 144
	XMR        = ZeroQuote + 128
	NEAR       = ZeroQuote + 397
	DOT        = ZeroQuote + 354
	BSV        = ZeroQuote + 236
	AVAX       = ZeroQuote + 9000
	FTM        = ZeroQuote + 1007
)

// eth chain id https://chainlist.org/zh
// 1为ETH链，56为BSC链，61为ETC链，128为HECO链，137为Matic链, 250 FTM,
// 25 CRONOS, 42262 - Emerald, 42220 - CELO, 66 - OKXChain, 888 - WanChain , 10 - Optimism

// init net params
var (
	BTCParams        = chaincfg.MainNetParams
	BTCTestnetParams = chaincfg.TestNet3Params
	LTCParams        = chaincfg.MainNetParams
	DOGEParams       = chaincfg.MainNetParams
	DASHParams       = chaincfg.MainNetParams
	BCHParams        = chaincfg.MainNetParams
	QTUMParams       = chaincfg.MainNetParams
	USDTParams       = chaincfg.MainNetParams
	BSVParams        = chaincfg.MainNetParams
)

func init() {
	// ltc net params
	// https://github.com/litecoin-project/litecoin/blob/master/src/chainparams.cpp
	LTCParams.Bech32HRPSegwit = "ltc"
	LTCParams.PubKeyHashAddrID = 0x30 // 48
	LTCParams.ScriptHashAddrID = 0x32 // 50
	LTCParams.PrivateKeyID = 0xb0     // 176

	// doge net params
	// https://github.com/dogecoin/dogecoin/blob/master/src/chainparams.cpp
	DOGEParams.PubKeyHashAddrID = 0x1e // 30
	DOGEParams.ScriptHashAddrID = 0x16 // 22
	DOGEParams.PrivateKeyID = 0x9e     // 158

	// dash net params
	// https://github.com/dashpay/dash/blob/master/src/chainparams.cpp
	DASHParams.PubKeyHashAddrID = 0x4c // 76
	DASHParams.ScriptHashAddrID = 0x10 // 16
	DASHParams.PrivateKeyID = 0xcc     // 204

	// bch net params
	// https://github.com/Bitcoin-ABC/bitcoin-abc/blob/master/src/chainparams.cpp
	BCHParams.PubKeyHashAddrID = 0x00 // 0
	BCHParams.ScriptHashAddrID = 0x05 // 5
	BCHParams.PrivateKeyID = 0x80     // 128

	// qtum net params
	// https://github.com/qtumproject/qtum/blob/master/src/chainparams.cpp
	QTUMParams.PubKeyHashAddrID = 0x3a // 58
	QTUMParams.ScriptHashAddrID = 0x32 // 50
	QTUMParams.PrivateKeyID = 0x80     // 128

	// usdt net params
	// https://github.com/OmniLayer/omnicore/blob/master/src/chainparams.cpp
	USDTParams.PubKeyHashAddrID = 0x00 // 0
	USDTParams.ScriptHashAddrID = 0x05 // 5
	USDTParams.PrivateKeyID = 0x80     // 128

	// bsv net params
	// https://github.com/bitcoin-sv/bitcoin-sv/blob/master/src/chainparams.cpp
	BSVParams.PubKeyHashAddrID = 0x00 // 0
	BSVParams.ScriptHashAddrID = 0x05 // 5
	BSVParams.PrivateKeyID = 0x80     // 128
}
