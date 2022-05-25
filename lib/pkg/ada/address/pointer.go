package address

import (
	"wallet-srv/lib/pkg/ada/network"

	"github.com/btcsuite/btcutil/bech32"
	"github.com/fxamacker/cbor/v2"
)

type PointerAddress struct {
	Network network.NetworkInfo
}

func (p *PointerAddress) MarshalCBOR() (bytes []byte, err error) {
	return cbor.Marshal(p.Bytes())
}

func (p *PointerAddress) Bytes() (bytes []byte) {
	return
}

func (p *PointerAddress) String() string {
	str, _ := bech32.Encode(p.Prefix(), p.Bytes())
	return str
}

func (p *PointerAddress) Prefix() string {
	if p.Network == *network.TestNet() {
		return "addr_test"
	}
	return "addr"
}

func (p *PointerAddress) NetworkInfo() *network.NetworkInfo {
	return &(p.Network)
}
