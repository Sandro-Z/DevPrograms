package ecdsa

import (
	"crypto/ecdsa"
	"crypto/rand"

	"git.ana/xjtuana/deland"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	ecdsax "github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
)

func GeneratePrivateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(btcec.S256(), rand.Reader)
}

func Recover(signature []byte, hash deland.Hash) (*ecdsa.PublicKey, error) {
	key, _, err := ecdsax.RecoverCompact(signature[:], hash.Bytes())
	return key.ToECDSA(), err
}

func Sign(prv *ecdsa.PrivateKey, hash deland.Hash) ([]byte, error) {
	key := secp256k1.PrivKeyFromBytes(prv.D.Bytes())
	return ecdsax.SignCompact(key, hash.Bytes(), false), nil
}

func Verify(signature []byte, hash deland.Hash, pub *ecdsa.PublicKey) bool {
	if len(signature) != 65 {
		return false
	}
	sigRecoveryCode := signature[0]
	_ = sigRecoveryCode
	var r, s secp256k1.ModNScalar
	if r.SetByteSlice(signature[1:33]) || r.IsZero() {
		return false
	}
	if s.SetByteSlice(signature[33:]) || s.IsZero() {
		return false
	}
	var x, y secp256k1.FieldVal
	if x.SetByteSlice(pub.X.Bytes()) || r.IsZero() {
		return false
	}
	if y.SetByteSlice(pub.Y.Bytes()) || s.IsZero() {
		return false
	}
	return ecdsax.NewSignature(&r, &s).Verify(hash.Bytes(), secp256k1.NewPublicKey(&x, &y))
}
