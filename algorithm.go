package otpgo

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

type hmacAlgorithm int

const (
	// To use the SHA1 hash algorithm to calculate an HMAC.
	HmacSHA1 hmacAlgorithm = iota + 1
	// To use the SHA256 hash algorithm to calculate an HMAC.
	HmacSHA256
	// To use the SHA512 hash algorithm to calculate an HMAC.
	HmacSHA512
)

// Hash returns a hash.Hash instance corresponding to the hmacAlgorithm type.
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
