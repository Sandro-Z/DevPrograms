//go:generate protoc --go_out=$GOPATH/src pb/deland.proto

package deland

import (
	"encoding/hex"
	"math/big"
)

const (
	AddressLength = 20
	HashLength    = 32
)

type Address [AddressLength]byte

func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}

func (a Address) Bytes() []byte { return a[:] }
func (a Address) Big() *big.Int { return new(big.Int).SetBytes(a[:]) }
func (a Address) Hex() string   { return hex.EncodeToString(a[:]) }

func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
}

type Bytes []byte

func (b Bytes) Hex() string { return hex.EncodeToString(b[:]) }

type Hash [HashLength]byte

func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

func (h Hash) Bytes() []byte { return h[:] }
func (h Hash) Big() *big.Int { return new(big.Int).SetBytes(h[:]) }
func (h Hash) Hex() string   { return hex.EncodeToString(h[:]) }

func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashLength:]
	}
	copy(h[HashLength-len(b):], b)
}

type Record struct {
	From *Address `json:"from"`
	Data *Bytes   `json:"data"`

	V *Bytes `json:"v"`
	R *Bytes `json:"r"`
	S *Bytes `json:"s"`
}
