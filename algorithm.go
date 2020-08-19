package otpgo

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

type hmacAlgorithm int

const (
	HmacSHA1 hmacAlgorithm = iota + 1
	HmacSHA256
	HmacSHA512
)

// Returns a hash.Hash instance corresponding to the hmacAlgorithm type.
func (alg hmacAlgorithm) Hash() (h hash.Hash) {
	switch alg {
	case HmacSHA1:
		h = sha1.New()
	case HmacSHA256:
		h = sha256.New()
	case HmacSHA512:
		h = sha512.New()
	default:
		panic("unexpected hash algorithm")
	}

	return h
}
