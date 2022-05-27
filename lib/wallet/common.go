package wallet

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"wallet-srv/model"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
	"github.com/ethereum/go-ethereum/params"
	"github.com/tyler-smith/go-bip39"
)

type SegWitType int

const (
	SymbolEth  = "ETH"
	SymbolBtc  = "BTC"
	SymbolTrx  = "TRX"
	SymbolBch  = "BCH"
	SymbolLtc  = "LTC"
	SymbolDash = "DASH"
	SymbolDoge = "DOGE"
	SymbolQTUM = "QTUM"
	SymbolSol  = "SOL"
	SymbolFil  = "FIL"
	SymbolLuna = "LUNA"
	SymbolAda  = "ADA"
	SymbolXrp  = "XRP"
	SymbolXmr  = "XMR"
	SymbolNear = "NEAR"
	SymbolDot  = "DOT"
	SymbolBSV  = "BSV"
	SymbolAvax = "AVAX"
	SymbolFlow = "FLOW"

	BtcChainMainNet  = int(wire.MainNet)
	BtcChainTestNet3 = int(wire.TestNet3)
	BtcChainRegtest  = int(wire.TestNet)
	BtcChainSimNet   = int(wire.SimNet)

	//chainId
	EthChainId = 1

	SegWitNone   SegWitType = 0
	SegWitScript SegWitType = 1
	SegWitNative SegWitType = 2

	ChangeTypeExternal = 0
	ChangeTypeInternal = 1 // Usually used for change, not visible to the outside world

	SatoshiPerBitcoin = 1e8
	SunPerTrx         = 1e6
	GweiPerEther      = 1e9
	WeiPerGwei        = 1e9
	WeiPerEther       = 1e18

	EtherTransferGas = 21000

	TokenShowDecimals = 9
)

var IsFixIssue172 = false

func NewEntropy(bits int) (entropy []byte, err error) {
	return bip39.NewEntropy(bits)
}

func NewMnemonic(bits int) (mnemonic string, err error) {
	entropy, err := NewEntropy(bits)
	if err != nil {
		return "", err
	}
	return NewMnemonicByEntropy(entropy)
}

func NewSeed(len uint8) ([]byte, error) {
	return hdkeychain.GenerateSeed(len)
}

func NewMnemonicByEntropy(entropy []byte) (mnemonic string, err error) {
	return bip39.NewMnemonic(entropy)
}

func EntropyFromMnemonic(mnemonic string) (entropy []byte, err error) {
	return bip39.EntropyFromMnemonic(mnemonic)
}

func NewSeedFromMnemonic(mnemonic, password string) ([]byte, error) {
	if mnemonic == "" {
		return nil, errors.New("mnemonic is required")
	}
	return bip39.NewSeedWithErrorChecking(mnemonic, password)
}

func MakeBip44Path(coinType uint32, accountIndex, changeType, index int) (string, error) {
	return MakeBipXPath(44, coinType, accountIndex, changeType, index)
}

func MakeBip49Path(coinType uint32, accountIndex, changeType, index int) (string, error) {
	return MakeBipXPath(49, coinType, accountIndex, changeType, index)
}

func MakeBip84Path(coinType uint32, accountIndex, changeType, index int) (string, error) {
	return MakeBipXPath(84, coinType, accountIndex, changeType, index)
}

func MakeBipXPath(bipType int, coinType uint32, accountIndex, changeType, index int) (string, error) {
	if accountIndex < 0 || index < 0 {
		return "", errors.New("invalid account index or index")
	}
	if changeType != ChangeTypeExternal && changeType != ChangeTypeInternal {
		return "", errors.New("invalid change type")
	}
	return fmt.Sprintf("m/%d'/%d'/%d'/%d/%d", bipType, coinType, accountIndex, changeType, index), nil
}

func FormatBtc(amount int64) string {
	return FormatFloat(float64(amount)/SatoshiPerBitcoin, 8)
}

func FormatEth(amount int64) string {
	return FormatFloat(float64(amount)/GweiPerEther, 9)
}

func FormatFloat(f float64, precision int) string {
	d := float64(1)
	if precision > 0 {
		d = math.Pow10(precision)
	}
	return strconv.FormatFloat(math.Trunc(f*d)/d, 'f', -1, 64)
}

func GetChainParam(symbol string) *chaincfg.Params {
	var chainParam *chaincfg.Params
	switch strings.ToUpper(symbol) {
	case "BTC":
		chainParam = &BTCParams
	case "BCH":
		chainParam = &BCHParams
	case "LTC":
		chainParam = &LTCParams
	case "DOGE":
		chainParam = &DOGEParams
	case "DASH":
		chainParam = &DASHParams
	case "QTUM":
		chainParam = &QTUMParams
	case "BSV":
		chainParam = &BSVParams
	default:
		chainParam = &chaincfg.MainNetParams
	}
	return chainParam
}

func GetChainConfig(symbol string) *params.ChainConfig {

	var chainConfig *params.ChainConfig
	switch symbol {
	case "ETH":
		chainConfig = &params.ChainConfig{}
	default:
		chainConfig = &params.ChainConfig{}
	}
	return chainConfig
}

func GetDBModel(symbol string) (model.IfAddress, string) {

	symbol = strings.ToLower(symbol)

	var AddrModelMap = map[string]model.IfAddress{
		"btc":  &model.BtcAddress{},
		"bch":  &model.BchAddress{},
		"doge": &model.DogeAddress{},
		"dash": &model.DashAddress{},
		"ltc":  &model.LtcAddress{},
		"qtum": &model.QtumAddress{},
		"eth":  &model.EthAddress{},
		"trx":  &model.TrxAddress{},
		"sol":  &model.SolAddress{},
		"fil":  &model.FilAddress{},
		"ada":  &model.AdaAddress{},
		"xrp":  &model.XrpAddress{},
		"luna": &model.LunaAddress{},
		"xmr":  &model.XmrAddress{},
		"near": &model.NearAddress{},
		"dot":  &model.DotAddress{},
		"bsv":  &model.BsvAddress{},
		"avax": &model.AvaxAddress{},
		"flow": &model.FlowAddress{},
	}

	return AddrModelMap[symbol], symbol + "_address"
}

func GetCoinType(symbol string) uint32 {

	coinType := BTC
	switch strings.ToLower(symbol) {
	case "bch":
		coinType = BCH
	case "dash":
		coinType = DASH
	case "doge":
		coinType = DOGE
	case "ltc":
		coinType = LTC
	case "qtum":
		coinType = QTUM
	case "bsv":
		coinType = BSV
	default:
		coinType = BTC
	}
	return coinType
}
