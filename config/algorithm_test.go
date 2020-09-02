package config

import (
	"testing"
)

func TestHmacAlgorithm_Hash(t *testing.T) {
	cases := []struct {
		label             string
		alg               HmacAlgorithm
		expectedSize      int
		expectedBlockSize int
		shouldPanic       bool
	}{
		{label: "HmacSHA1", alg: HmacSHA1, expectedSize: 20, expectedBlockSize: 64},
		{label: "HmacSHA256", alg: HmacSHA256, expectedSize: 32, expectedBlockSize: 64},
		{label: "HmacSHA512", alg: HmacSHA512, expectedSize: 64, expectedBlockSize: 128},
		{label: "Panicky", alg: HmacAlgorithm(-1), shouldPanic: true},
	}

	assertPanic := func(t *testing.T, label string, alg HmacAlgorithm) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("case %s: was expected to panic", label)
			}
		}()
		alg.Hash()
	}

	for _, c := range cases {
		if c.shouldPanic {
			assertPanic(t, c.label, c.alg)
			break
		}

		h := c.alg.Hash()

		if c.expectedSize != h.Size() {
			t.Errorf("case %s: wrong hash size\nexpected: %d\n  actual: %d", c.label, c.expectedSize, h.Size())
			t.FailNow()
		}

		if c.expectedBlockSize != h.BlockSize() {
			t.Errorf(
				"case %s: wrong hash block size\nexpected: %d\n  actual: %d",
				c.label,
				c.expectedBlockSize,
				h.BlockSize(),
			)
			t.FailNow()
		}
	}
}

func TestHmacAlgorithm_String(t *testing.T) {
	cases := []struct {
		label            string
		alg              HmacAlgorithm
		expectedReadable string
		shouldPanic      bool
	}{
		{label: "HmacSHA1", alg: HmacSHA1, expectedReadable: "SHA1"},
		{label: "HmacSHA256", alg: HmacSHA256, expectedReadable: "SHA256"},
		{label: "HmacSHA512", alg: HmacSHA512, expectedReadable: "SHA512"},
		{label: "Panicky", alg: HmacAlgorithm(-1), shouldPanic: true},
	}

	assertPanic := func(t *testing.T, label string, alg HmacAlgorithm) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("case %s: was expected to panic", label)
			}
		}()
		_ = alg.String()
	}

	for _, c := range cases {
		if c.shouldPanic {
			assertPanic(t, c.label, c.alg)
			break
		}

		h := c.alg.String()

		if c.expectedReadable != h {
			t.Errorf("case %s: wrong hash readable string\nexpected: %s\n  actual: %s", c.label, c.expectedReadable, h)
		}
	}
}
