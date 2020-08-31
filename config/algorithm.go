package config

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

type HmacAlgorithm int

const (
	// HmacSHA1 indicates to use the SHA1 hash to calculate an HMAC.
	HmacSHA1 HmacAlgorithm = iota + 1
	// HmacSHA256 indicates to use the SHA256 hash to calculate an HMAC.
	HmacSHA256
	// HmacSHA512 indicates to use the SHA512 hash to calculate an HMAC.
	HmacSHA512
)

// Hash returns a hash.Hash instance corresponding to the HmacAlgorithm type.
func (alg HmacAlgorithm) Hash() (h hash.Hash) {
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

func (alg HmacAlgorithm) String() string {
	var s string
	switch alg {
	case HmacSHA1:
		s = "SHA1"
	case HmacSHA256:
		s = "SHA256"
	case HmacSHA512:
		s = "SHA512"
	default:
		panic("unexpected hash algorithm")
	}

	return s
}
