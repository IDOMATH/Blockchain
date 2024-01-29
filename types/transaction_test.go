package types

import (
	"github.com/idomath/Blockchain/crypto"
	"github.com/idomath/Blockchain/proto"
	"github.com/idomath/Blockchain/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTransaction(t *testing.T) {
	fromPrivateKey := crypto.GeneratePrivateKey()
	fromAddress := fromPrivateKey.Public().Address().Bytes()

	toPrivateKey := crypto.GeneratePrivateKey()
	toAddress := toPrivateKey.Public().Address().Bytes()

	input := &proto.TxInput{
		PrevTxHash:   util.RandomHash(),
		PrevOutIndex: 0,
		PublicKey:    fromPrivateKey.Public().Bytes(),
	}

	output1 := &proto.TxOutput{
		Amount:  5,
		Address: toAddress,
	}
	output2 := &proto.TxOutput{
		Amount:  95,
		Address: fromAddress,
	}

	tx := &proto.Transaction{
		Version: 1,
		Inputs:  []*proto.TxInput{input},
		Outputs: []*proto.TxOutput{output1, output2},
	}
	sig := SignTransaction(fromPrivateKey, tx)
	input.Signature = sig.Bytes()

	assert.True(t, VerifyTransactino(tx))
}
