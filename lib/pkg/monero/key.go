package monero

import (
	"crypto/rand"
	"io"
	"unsafe"

	"golang.org/x/crypto/sha3"
)

const (
	KeyLength = 32
)

// Key can be a Scalar or a Point
type Key [KeyLength]byte

func (p *Key) FromBytes(b [KeyLength]byte) {
	*p = b
}

func (p *Key) ToBytes() (result [KeyLength]byte) {
	result = [KeyLength]byte(*p)
	return
}

func (p *Key) PubKey() (pubKey *Key) {
	point := new(ExtendedGroupElement)
	GeScalarMultBase(point, p)
	pubKey = new(Key)
	point.ToBytes(pubKey)
	return
}

// Creates a point on the Edwards Curve by hashing the key
func (p *Key) HashToEC() (result *ExtendedGroupElement) {
	result = new(ExtendedGroupElement)
	var p1 ProjectiveGroupElement
	var p2 CompletedGroupElement
	h := Key(Keccak256(p[:]))
	p1.FromBytes(&h)
	GeMul8(&p2, &p1)
	p2.ToExtended(result)
	return
}

func (p *Key) Address() string {

	// The Private Spend Key and Private View Key are sent to the ed25519 scalarmult function
	// to create their counterparts, the Public Spend Key and Public View Key
	spendKey := p.ToBytes()
	viewKey := p.PubKey().ToBytes()
	spendPub := publicKeyFromPrivateKey(&spendKey)
	viewPub := publicKeyFromPrivateKey(&viewKey)

	// monero main network
	network := []byte{0x12}

	// The pair of public keys are prepended with one network byte (the number 18, 0x12, for Monero).
	// It looks like this: (network byte) + (32-byte public spend key) + (32-byte public view key).
	// These 65 bytes are hashed with Keccak-256.
	hash := keccak256(network, spendPub[:], viewPub[:])

	// The first four bytes of the hash are appended, creating a 69-byte Public Address.
	// As a last step, this 69-byte string is converted to Base58.
	// However, it's not done all at once like a Bitcoin address, but rather in 8-byte blocks.
	// This gives us eight full-sized blocks and one 5-byte block. Eight bytes converts to 11 or less Base58 characters;
	// if a particular block converts to <11 characters, the conversion pads it with "1"s (1 is 0 in Base58).
	// Likewise, the final 5-byte block can convert to 7 or less Base58 digits; the conversion will ensure the result is 7 digits.
	// Due to the conditional padding, the 69-byte string will always convert to 95 Base58 characters (8 * 11 + 7).
	address := EncodeMoneroBase58(network, spendPub[:], viewPub[:], hash[:4])

	return address
}

func RandomScalar() (result *Key) {
	result = new(Key)
	var reduceFrom [KeyLength * 2]byte
	tmp := make([]byte, KeyLength*2)
	rand.Read(tmp)
	copy(reduceFrom[:], tmp)
	ScReduce(result, &reduceFrom)
	return
}

func NewKeyPair() (privKey *Key, pubKey *Key) {
	privKey = RandomScalar()
	pubKey = privKey.PubKey()
	return
}

func ParseKey(buf io.Reader) (result Key, err error) {
	key := make([]byte, KeyLength)
	if _, err = buf.Read(key); err != nil {
		return
	}
	copy(result[:], key)
	return
}

func NewKey(seed []byte) (result *Key) {
	result = new(Key)
	var reduceFrom [KeyLength * 2]byte
	copy(reduceFrom[:], seed)
	ScReduce(result, &reduceFrom)
	return
}

func keccak256(data ...[]byte) *[KeyLength]byte {
	h := sha3.NewLegacyKeccak256()
	for _, v := range data {
		h.Write(v)
	}
	sum := h.Sum(nil)
	sum32 := (*[KeyLength]byte)(unsafe.Pointer(&sum[0]))

	return sum32
}

func publicKeyFromPrivateKey(priv *[KeyLength]byte) *[KeyLength]byte {
	pub := new([KeyLength]byte)

	p := new(ExtendedGroupElement)
	GeScalarMultBase(p, (*Key)(priv))
	p.ToBytes((*Key)(pub))

	return pub
}
