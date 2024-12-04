package main

import (
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"log"

	"git.ana/xjtuana/deland/pkg/crypto"
	"git.ana/xjtuana/deland/pkg/crypto/ecdsa"
	"github.com/btcsuite/btcd/btcec/v2"
)

func main() {
	log.Println("===== Client ECDSA =====")
	privateKey, err := ecdsa.GeneratePrivateKey()
	if err != nil {
		log.Panicln(err)
	}
	log.Println("PrivateKey: 0x" + hex.EncodeToString(privateKey.D.Bytes()))
	log.Println("PublicKey : 0x" + hex.EncodeToString(elliptic.Marshal(btcec.S256(), privateKey.X, privateKey.Y)))
	address := crypto.PublicKeyToAddress(privateKey.PublicKey)
	log.Println("Address   : 0x" + hex.EncodeToString(address.Bytes()))

	log.Println("===== Sign =====")
	digest := sha256.Sum256([]byte("hello"))
	log.Println("Digest    : 0x" + hex.EncodeToString(digest[:]))
	signature, err := ecdsa.Sign(privateKey, digest)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Signature : 0x" + hex.EncodeToString(signature))

	log.Println("===== Recover =====")
	publicKey, err := ecdsa.Recover(signature, digest)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("PublicKey : 0x" + hex.EncodeToString(elliptic.Marshal(btcec.S256(), publicKey.X, publicKey.Y)))

	log.Println("===== Verify =====")
	log.Println("Verify    :", ecdsa.Verify(signature, digest, publicKey))
}
