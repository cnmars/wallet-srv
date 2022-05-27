package wallet

import (
	"encoding/hex"
	"fmt"
	"wallet-srv/lib/pkg/secp256k1"

	"github.com/ava-labs/avalanchego/utils/crypto"
)

type HDWallet struct {
	seed     []byte
	symbol   string
	coinType uint32
}

func NewHDWallet(symbol string, coinType uint32) (*HDWallet, error) {

	var seed []byte
	var err error
	switch symbol {
	case SymbolFil:
		seed, err = secp256k1.GenerateKey()
	case SymbolXmr:
		seed, err = NewSeed(64)
	case SymbolAvax:
		factory := crypto.FactorySECP256K1R{}
		key, _ := factory.NewPrivateKey()
		seed = key.Bytes()
	default:
		seed, err = NewSeed(32)
	}

	if err != nil {
		return nil, err
	}

	return &HDWallet{
		seed:     seed,
		symbol:   symbol,
		coinType: coinType}, nil
}

func NewWalletSeed(symbol string, coinType uint32, s string) (*HDWallet, error) {

	seed, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return &HDWallet{seed: seed, symbol: symbol, coinType: coinType}, nil
}

func (hd *HDWallet) NewWallet(accountIndex, changeType, index int) (Wallet, error) {
	path, err := MakeBip44Path(hd.coinType, accountIndex, changeType, index)
	if err != nil {
		return nil, err
	}

	return hd.NewWalletByPath(hd.symbol, hd.coinType, path, SegWitNone)
}

func (hd *HDWallet) NewSegWitWallet(accountIndex, changeType, index int) (Wallet, error) {
	path, err := MakeBip49Path(hd.coinType, accountIndex, changeType, index)
	if err != nil {
		return nil, err
	}
	return hd.NewWalletByPath(hd.symbol, hd.coinType, path, SegWitScript)
}

func (hd *HDWallet) NewNativeSegWitWallet(accountIndex, changeType, index int) (Wallet, error) {
	path, err := MakeBip84Path(hd.coinType, accountIndex, changeType, index)
	if err != nil {
		return nil, err
	}
	return hd.NewWalletByPath(hd.symbol, hd.coinType, path, SegWitNative)
}

func (hd *HDWallet) NewWalletByPath(symbol string, coinType uint32, path string, segWitType SegWitType) (Wallet, error) {
	var w Wallet
	var err error

	switch symbol {
	case "BTC", "BCH", "LTC", "DASH", "DOGE", "QTUM", "BSV":
		params := GetChainParam(symbol)
		w, err = NewBtcWalletByPath(symbol, path, hd.seed, segWitType, params)
	case "ETH":
		w, err = NewEthWalletByPath(path, hd.seed, EthChainId)
	case "TRX":
		w, err = NewTrxWalletByPath(path, hd.seed)
	case "SOL":
		w, err = NewSolWallet(hd.seed)
	case "FIL":
		w, err = NewFilWallet(hd.seed)
	case "ADA":
		w, err = NewAdaWallet(hd.seed)
	case "XRP":
		w, err = NewXrpWallet(hd.seed)
	case "LUNA":
		w, err = NewLunaWallet(hd.seed)
	case "XMR":
		w, err = NewXmrWallet(hd.seed)
	case "NEAR":
		w, err = NewNearWallet(hd.seed)
	case "DOT":
		w, err = NewDotWallet(hd.seed)
	case "AVAX":
		w, err = NewAvaxWallet(hd.seed)
	case "FLOW":
		w, err = NewFlowWallet(hd.seed)
	default:
		err = fmt.Errorf("invalid symbol: %s", symbol)
	}

	if err != nil {
		return nil, err
	}
	return w, nil
}

func (hd *HDWallet) GetSeed() string {
	return hex.EncodeToString(hd.seed)
}
