package util

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
)

type Hash []byte

type Hasher func([]byte) Hash

type Hashable interface {
	Hash(Hasher) Hash
}

type RLPHashable interface {
	RLPHash(Hasher) Hash
}

func DoHash(b []byte) Hash {
	hash := sha3.NewLegacyKeccak256()

	var buf []byte
	hash.Write(b)
	buf = hash.Sum(buf)

	return buf
}

func Sha256(b []byte) Hash {
	hash := sha3.Sum256(b)
	return hash[:]
}

func ValidateSignature(hash, signature []byte, address common.Address) error {
	if len(signature) == 65 && signature[64] > 26 {
		signature[64] -= 27
	}
	pubKey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return err
	}
	signatureAddress := crypto.PubkeyToAddress(*pubKey)
	if !AddressesEqual(&address, &signatureAddress) {
		return errors.New("Invalid signature")
	}
	return nil
}
