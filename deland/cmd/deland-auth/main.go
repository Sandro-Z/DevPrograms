package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"log"
	"time"

	"filippo.io/edwards25519"
	"git.ana/xjtuana/deland/pkg/crypto"
	"git.ana/xjtuana/deland/pkg/crypto/ecdsa"
	"git.ana/xjtuana/deland/pkg/crypto/ed25519"
	"github.com/btcsuite/btcd/btcec/v2"
)

var (
	version = []byte(time.Now().UTC().Format(time.RFC3339))
	domain  = []byte("example.com")
)

var (
	serverPublicKey, serverPrivateKey, _ = ed25519.GenerateKey(rand.Reader)
)

func init() {
	log.Println("===== Server ED25519 =====")
	log.Println("Version   : " + string(version))
	log.Println("PrivateKey: 0x" + hex.EncodeToString(serverPrivateKey))
}

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

	log.Println("===== Nonce =====")
	RR, r, err := ed25519.Nonce(serverPrivateKey)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("R'        : 0x" + hex.EncodeToString(RR.Bytes()))

	log.Println("===== Blind =====")
	ra, rb := make([]byte, 64), make([]byte, 64)
	rand.Read(ra)
	rand.Read(rb)
	a, err := edwards25519.NewScalar().SetUniformBytes(ra)
	if err != nil {
		log.Panicln(err)
	}
	b, err := edwards25519.NewScalar().SetUniformBytes(rb)
	if err != nil {
		log.Panicln(err)
	}
	R, mm, err := ed25519.Blind(serverPublicKey, RR, a, b, address[:])
	if err != nil {
		log.Panicln(err)
	}
	log.Println("m'        : 0x" + hex.EncodeToString(mm.Bytes()))

	log.Println("===== Sign =====")
	ss, err := ed25519.Sign(serverPrivateKey, r, version, domain, mm)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("s'        : 0x" + hex.EncodeToString(ss.Bytes()))

	log.Println("===== Unblind =====")
	signature := ed25519.Unblind(R, ss, a)
	log.Println("Signature : 0x" + hex.EncodeToString(signature))

	log.Println("===== Verify =====")
	log.Println("Verify    :", ed25519.Verify(signature, version, domain, address[:], serverPublicKey))
}
