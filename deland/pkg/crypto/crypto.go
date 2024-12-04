package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"

	"git.ana/xjtuana/deland"
	"github.com/btcsuite/btcd/btcec/v2"
)

func PublicKeyToAddress(pub ecdsa.PublicKey) deland.Address {
	digest := sha256.Sum256(elliptic.Marshal(btcec.S256(), pub.X, pub.Y)[1:])
	return deland.BytesToAddress(digest[12:])
}
