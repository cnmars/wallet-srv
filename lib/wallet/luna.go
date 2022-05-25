package wallet

import (
	"wallet-srv/lib/pkg/terra/msg"
	core "wallet-srv/lib/pkg/terra/types"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func init() {
	sdkConfig := sdk.GetConfig()
	sdkConfig.SetCoinType(core.CoinType)
	sdkConfig.SetFullFundraiserPath(core.FullFundraiserPath)
	sdkConfig.SetBech32PrefixForAccount(core.Bech32PrefixAccAddr, core.Bech32PrefixAccPub)
	sdkConfig.SetBech32PrefixForValidator(core.Bech32PrefixValAddr, core.Bech32PrefixValPub)
	sdkConfig.SetBech32PrefixForConsensusNode(core.Bech32PrefixConsAddr, core.Bech32PrefixConsPub)
	sdkConfig.SetAddressVerifier(core.AddressVerifier)
	sdkConfig.Seal()
}

type LunaWallet struct {
	symbol     string
	privateKey types.PrivKey
	publicKey  types.PubKey
}

func NewLunaWallet(seed []byte) (*LunaWallet, error) {

	algo, err := keyring.NewSigningAlgoFromString(string(hd.Secp256k1Type), keyring.SigningAlgoList{hd.Secp256k1})
	if err != nil {
		return nil, err
	}
	privKey := algo.Generate()(seed)
	pubkey := privKey.PubKey()
	return &LunaWallet{
		symbol:     SymbolLuna,
		privateKey: privKey,
		publicKey:  pubkey}, nil
}

func (w *LunaWallet) ChainId() int {
	return 0
}

func (w *LunaWallet) Symbol() string {
	return w.symbol
}

func (w *LunaWallet) DeriveAddress() string {
	addr := msg.AccAddress(w.privateKey.PubKey().Address())
	return addr.String()
}

func (w *LunaWallet) DerivePublicKey() string {
	return w.publicKey.String()
}

func (w *LunaWallet) DerivePrivateKey() string {
	return w.privateKey.String()
}

func (w *LunaWallet) DeriveNativePrivateKey() types.PrivKey {
	return w.privateKey
}
