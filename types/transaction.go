package types

import (
	"crypto/sha256"
	pb "github.com/golang/protobuf/proto"
	"github.com/idomath/Blockchain/crypto"
	"github.com/idomath/Blockchain/proto"
)

func SignTransaction(pk *crypto.PrivateKey, tx *proto.Transaction) *crypto.Signature {
	return pk.Sign(HashTransaction(tx))
}

func HashTransaction(tx *proto.Transaction) []byte {
	b, err := pb.Marshal(tx)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(b)
	return hash[:]
}

func VerifyTransactino(tx *proto.Transaction) bool {
	for _, input := range tx.Inputs {
		sig := crypto.SignatureFromBytes(input.Signature)
		publicKey := crypto.PublicKeyFromBytes(input.PublicKey)
		input.Signature = nil
		if !sig.Verify(publicKey, HashTransaction(tx)) {
			return false
		}
	}
	return true
}
