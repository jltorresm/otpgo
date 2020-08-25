package otpgo

import (
	"math"
	"time"
)

const (
	// TOTPDefaultPeriod is the default time period to use if none is provided by the caller.
	TOTPDefaultPeriod = 30
)

// The TOTP type used to generate Time-Based One-Time Passwords.
type TOTP struct {
	Key       string // Secret
	Period    int    // In seconds
	Algorithm hmacAlgorithm
	Length    otpLength
}

// Generate a Time-Based One-Time Password.
func (t *TOTP) Generate() (string, error) {
	// Make sure we have sensible values to generate secure OTPs
	t.ensureDefaults()

	// Make sure we have a valid non-empty key
	if err := t.ensureKey(); err != nil {
		return "", err
	}

	// Get the counter based on the current time
	timestamp := time.Now().Unix()
	counter := uint64(math.Floor(float64(timestamp) / float64(t.Period)))

	return generateOTP(t.Key, counter, t.Length, t.Algorithm)
}

// ensureDefaults applies sensible default values, if any of them is empty, so
// that the OTP generation works properly.
// Defaults:
//     - Period = TOTPDefaultPeriod = 30
//     - Algorithm = SHA1
//     - Length = 6
func (t *TOTP) ensureDefaults() {
	if t.Period == 0 {
		t.Period = TOTPDefaultPeriod
	}

	if t.Algorithm == 0 {
		t.Algorithm = HmacSHA1
	}

	if t.Length == 0 {
		t.Length = Length6
	}
}

// ensureKey generates a proper random key if no value is provided by the caller.
func (t *TOTP) ensureKey() (err error) {
	if t.Key != "" {
		return nil
	}

	t.Key, err = randomKey(RandomKeyLength)

	return err
}
