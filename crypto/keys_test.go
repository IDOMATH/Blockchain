package crypto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKey := GeneratePrivateKey()
	assert.Equal(t, len(privateKey.Bytes()), privateKeyLength)

	publicKey := privateKey.Public()
	assert.Equal(t, len(publicKey.Bytes()), publicKeyLength)
}

func TestPrivateKeySign(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.Public()
	msg := []byte("foo bar baz")

	sig := privateKey.Sign(msg)
	assert.True(t, sig.Verify(publicKey, msg))

	// Test with wrong message
	assert.False(t, sig.Verify(publicKey, []byte("foo")))

	//Test with wrong public key
	wrongPrivateKey := GeneratePrivateKey()
	wrongPublicKey := wrongPrivateKey.Public()
	assert.False(t, sig.Verify(wrongPublicKey, msg))
}

func TestPublicKeyToAddress(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.Public()
	address := publicKey.Address()
	assert.Equal(t, addressLength, len(address.Bytes()))

	fmt.Println(address)
}

func TestNewPrivateKeyFromString(t *testing.T) {
	var (
		seed       = "b3484ef6a36788214e9da7ed4e00005fdb8f60d6784f3df9e61da6fb9c3586f8"
		privateKey = NewPrivateKeyFromString(seed)
		addressStr = "9cefdccd6c88312824b0a47c112cf162b644535d"
	)
	assert.Equal(t, privateKeyLength, len(privateKey.Bytes()))

	address := privateKey.Public().Address()
	assert.Equal(t, addressStr, address.String())
}
