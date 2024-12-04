package ed25519

import (
	"bytes"
	"crypto/ed25519"
	"crypto/sha512"
	"fmt"
	"strconv"
	"time"

	"filippo.io/edwards25519"
)

// ===== BASIC =====
// =================
// ===== Nonce =====
// seed, P = prv[:32], prv[32:]
// h = SHA512(seed)
// p, prefix = h[:32], h[32:]
// r = SHA512(prefix||M)
// R = r * G

// ===== Sign =====
// k = SHA512(R||P||M)
// S = k * p + r mod l
// sig = R||S

// ===== Verify =====
// k = SHA512(sig[:32]||P||M)
// S = sig[32:]
// R = k * (-P) + S * G
// sig[:32] == R

// ===== Proof =====
// k = SHA512(R||P||M)
// P = p * G
// S = k * p + r mod l
// R = k * (-P) + S * G = r * G

// ===== BLIND =====
// =================
// ===== Nonce =====
// seed, P = prv[:32], prv[32:]
// h = SHA512(seed)
// p, prefix = h[:32], h[32:]
// r = SHA512(prefix||timestamp)
// R' = r * G

// ===== Blind =====
// R = R' + (a * G) + (b * P)
// k = SHA512(R||P||M)
// m' = k + b mod l

// ===== Sign =====
// d = SHA512(V||D)
// s' = m' * p + r - d mod l

// ===== Unblind =====
// S = s' + a mod l
// sig = R||S

// ===== Verify =====
// d = SHA512(V||D)
// S = sig[32:] + d mod l
// k = SHA512(sig[:32]||P||M)
// R = k * (-P) + S * G
// sig[:32] == R

// ===== Proof =====
// k = SHA512(R||P||M)
// P = p * G
// S = (((k + b) * p + r - d) + a) + d mod l
// R = k * (-P) + S * G = (b * P) + (r + a) * G

const (
	PublicKeySize  = ed25519.PublicKeySize
	PrivateKeySize = ed25519.PrivateKeySize
	SignatureSize  = ed25519.SignatureSize
	SeedSize       = ed25519.SeedSize
)

var GenerateKey = ed25519.GenerateKey

// ===== Blind =====
// R = R' + (a * G) + (b * P)
// k = SHA512(R||P||M)
// m' = k + b mod l
func Blind(publicKey ed25519.PublicKey, RR *edwards25519.Point, a, b *edwards25519.Scalar, message []byte) (*edwards25519.Point, *edwards25519.Scalar, error) {
	P, err := (&edwards25519.Point{}).SetBytes(publicKey)
	if err != nil {
		return nil, nil, err
	}

	RA := (&edwards25519.Point{}).ScalarBaseMult(a)
	RB := (&edwards25519.Point{}).ScalarMult(b, P)
	R := (&edwards25519.Point{}).Add(RR, (&edwards25519.Point{}).Add(RA, RB))

	kh := sha512.New()
	kh.Write(R.Bytes())
	kh.Write(publicKey)
	kh.Write(message)
	kd := make([]byte, 0, sha512.Size)
	kd = kh.Sum(kd)
	k, err := edwards25519.NewScalar().SetUniformBytes(kd)
	if err != nil {
		return nil, nil, err
	}

	mm := edwards25519.NewScalar().Add(k, b)

	return R, mm, nil
}

// ===== Unblind =====
// S = s' + a mod l
// sig = R||S
func Unblind(R *edwards25519.Point, ss, a *edwards25519.Scalar) []byte {
	S := edwards25519.NewScalar().Add(ss, a)
	return append(R.Bytes(), S.Bytes()...)
}

// ===== Nonce =====
// seed, P = prv[:32], prv[32:]
// h = SHA512(seed)
// p, prefix = h[:32], h[32:]
// r = SHA512(prefix||timestamp)
// R' = r * G
func Nonce(privateKey ed25519.PrivateKey) (*edwards25519.Point, *edwards25519.Scalar, error) {
	if l := len(privateKey); l != PrivateKeySize {
		return nil, nil, fmt.Errorf("ed25519: bad private key length: " + strconv.Itoa(l))
	}

	h := sha512.Sum512(privateKey[:SeedSize])

	rh := sha512.New()
	rh.Write(h[32:])
	rh.Write([]byte(time.Now().UTC().Format(time.RFC3339Nano)))
	rd := make([]byte, 0, sha512.Size)
	rd = rh.Sum(rd)
	r, err := edwards25519.NewScalar().SetUniformBytes(rd)
	if err != nil {
		return nil, nil, err
	}

	RR := (&edwards25519.Point{}).ScalarBaseMult(r)

	return RR, r, nil
}

// ===== Sign =====
// d = SHA512(V||D)
// s' = m' * p + r - d mod l
func Sign(privateKey ed25519.PrivateKey, r *edwards25519.Scalar, version, domain []byte, mm *edwards25519.Scalar) (*edwards25519.Scalar, error) {
	h := sha512.Sum512(privateKey[:ed25519.SeedSize])
	p, err := edwards25519.NewScalar().SetBytesWithClamping(h[:32])
	if err != nil {
		return nil, err
	}

	dh := sha512.New()
	dh.Write(version)
	dh.Write(domain)
	dd := make([]byte, 0, sha512.Size)
	dd = dh.Sum(dd)
	d, err := edwards25519.NewScalar().SetUniformBytes(dd)
	if err != nil {
		return nil, err
	}

	ss := edwards25519.NewScalar().MultiplyAdd(mm, p, edwards25519.NewScalar().Subtract(r, d))

	return ss, nil
}

// ===== Verify =====
// d = SHA512(V||D)
// S = sig[32:] + d mod l
// k = SHA512(sig[:32]||P||M)
// R = k * (-P) + S * G
// sig[:32] == R
func Verify(signature, version, domain, message []byte, publicKey ed25519.PublicKey) bool {
	if l := len(publicKey); l != PublicKeySize {
		panic("ed25519: bad public key length: " + strconv.Itoa(l))
	}

	if len(signature) != SignatureSize || signature[63]&224 != 0 {
		return false
	}

	P, err := (&edwards25519.Point{}).SetBytes(publicKey)
	if err != nil {
		return false
	}

	s, err := edwards25519.NewScalar().SetCanonicalBytes(signature[32:])
	if err != nil {
		return false
	}

	dh := sha512.New()
	dh.Write(version)
	dh.Write(domain)
	dd := make([]byte, 0, sha512.Size)
	dd = dh.Sum(dd)
	d, err := edwards25519.NewScalar().SetUniformBytes(dd)
	if err != nil {
		return false
	}

	S := edwards25519.NewScalar().Add(s, d)

	kh := sha512.New()
	kh.Write(signature[:32])
	kh.Write(publicKey)
	kh.Write(message)
	kd := make([]byte, 0, sha512.Size)
	kd = kh.Sum(kd)
	k, err := edwards25519.NewScalar().SetUniformBytes(kd)
	if err != nil {
		return false
	}

	minusP := (&edwards25519.Point{}).Negate(P)
	R := (&edwards25519.Point{}).VarTimeDoubleScalarBaseMult(k, minusP, S)

	return bytes.Equal(signature[:32], R.Bytes())
}

// ===== Proof =====
// k = SHA512(R||P||M)
// P = p * G
// S = (((k + b) * p + r - d) + a) + d mod l
// R = k * (-P) + S * G = (b * P) + (r + a) * G
