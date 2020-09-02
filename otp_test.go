package otpgo

import (
	"testing"

	"github.com/jltorresm/otpgo/config"
)

func TestGenerateOTP(t *testing.T) {
	cases := []struct {
		label       string
		key         string
		expectedOtp string
	}{
		{"Pad", "4LOWUKIZK2YA=======", "551709"},
		{"No Pad", "4LOWUKIZK2YA", "551709"},
		{"Pad 2", "NAZXS===", "967352"},
		{"No Pad 2", "NAZXS", "967352"},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			otp, err := generateOTP(c.key, 1, config.Length6, config.HmacSHA1)

			if c.expectedOtp != otp {
				t.Errorf("unexpected otp\nexpected: %s\n  actual: %s", c.expectedOtp, otp)
			}

			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}
		})
	}
}

func TestRandomKey(t *testing.T) {
	cases := []struct {
		label          string
		length         uint
		expectedLength int
		expectedError  error
	}{
		{"Zero", 0, 0, nil},
		{"One", 1, 2, nil},
		{"Three", 3, 5, nil},
		{"Small", 10, 16, nil},
		{"Normal", 64, 103, nil},
		{"Big", 1024, 1639, nil},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			key, err := randomKey(c.length)

			if c.expectedLength != len(key) {
				t.Errorf("unexpected key length\nexpected: %d\n  actual: %d", c.expectedLength, len(key))
			}

			if c.expectedError != err {
				t.Errorf("unexpected error\nexpected: %s\n  actual: %s", c.expectedError, err)
			}

			if len(key) > 0 && key[len(key)-1:] == "=" {
				t.Errorf("unexpected padding, expected the key to have no padding")
			}
		})
	}
}
