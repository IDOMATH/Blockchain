package types

import (
	"github.com/idomath/Blockchain/crypto"
	"github.com/idomath/Blockchain/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignBlock(t *testing.T) {
	block := util.RandomBlock()
	privateKey := crypto.GeneratePrivateKey()
	publicKey := privateKey.Public()

	sig := SignBlock(privateKey, block)
	assert.Equal(t, 64, len(sig.Bytes()))
	assert.True(t, sig.Verify(publicKey, HashBlock(block)))
}

func TestHashBlock(t *testing.T) {
	block := util.RandomBlock()
	hash := HashBlock(block)
	assert.Equal(t, 32, len(hash))
}
