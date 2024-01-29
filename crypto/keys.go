package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"io"
)

const (
	privateKeyLength = 64
	signatureLength  = 64
	publicKeyLength  = 32
	seedLength       = 32
	addressLength    = 20
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

func NewPrivateKeyFromString(s string) *PrivateKey {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return NewPrivateKeyFromSeed(b)
}

func NewPrivateKeyFromSeed(seed []byte) *PrivateKey {
	if len(seed) != seedLength {
		panic("invalid seed length, must be 32")
	}

	return &PrivateKey{key: ed25519.NewKeyFromSeed(seed)}
}

func GeneratePrivateKey() *PrivateKey {
	seed := make([]byte, seedLength)
	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		panic(err)
	}

	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}

func (p *PrivateKey) Sign(msg []byte) *Signature {
	return &Signature{
		value: ed25519.Sign(p.key, msg),
	}
}

func (p *PrivateKey) Public() *PublicKey {
	b := make([]byte, publicKeyLength)
	copy(b, p.key[32:])
	return &PublicKey{key: b}
}

type PublicKey struct {
	key ed25519.PublicKey
}

func PublicKeyFromBytes(b []byte) *PublicKey {
	if len(b) != publicKeyLength {
		panic("invalid public key length")
	}
	return &PublicKey{
		key: ed25519.PublicKey(b),
	}
}

func (p *PublicKey) Address() Address {
	return Address{
		value: p.key[len(p.key)-addressLength:],
	}
}

func (p *PublicKey) Bytes() []byte {
	return p.key
}

type Signature struct {
	value []byte
}

func SignatureFromBytes(b []byte) *Signature {
	if len(b) != signatureLength {
		panic("length of bytes not equal to 64")
	}
	return &Signature{value: b}
}

func (s *Signature) Verify(publicKey *PublicKey, msg []byte) bool {
	return ed25519.Verify(publicKey.key, msg, s.value)
}

func (s *Signature) Bytes() []byte {
	return s.value
}

type Address struct {
	value []byte
}

func (a Address) Bytes() []byte {
	return a.value
}

func (a Address) String() string {
	return hex.EncodeToString(a.value)
}
