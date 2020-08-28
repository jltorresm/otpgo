package otpgo

import (
	"errors"
)

const (
	// HOTPDefaultLeeway is the default acceptable look-ahead look-behind window.
	// A value of 1 means the OTP will be valid if it coincides with the
	// calculated token for the current counter, the next one or the one before.
	HOTPDefaultLeeway uint64 = 1
)

// The HOTP type used to generate HMAC-Based One-Time Passwords.
type HOTP struct {
	Key       string // Secret
	Counter   uint64
	Leeway    uint64
	Algorithm hmacAlgorithm
	Length    otpLength
}

// Generate a HMAC-Based One-Time Password.
func (h *HOTP) Generate() (string, error) {
	// Make sure we have sensible values to generate secure OTPs
	h.ensureDefaults()

	// Make sure we have a valid non-empty key
	if err := h.ensureKey(); err != nil {
		return "", err
	}

	return generateOTP(h.Key, h.Counter, h.Length, h.Algorithm)
}

// Validate will try to check if the provided token is a valid OTP for the
// current HOTP config. If the validation is successful the internal Counter
// will be incremented by one.
func (h *HOTP) Validate(token string) (bool, error) {
	// Validating without a proper key shouldn't happen
	if h.Key == "" {
		return false, errors.New("missing secret key for validation")
	}

	// Make sure we have sensible values to generate secure OTPs
	h.ensureDefaults()

	// A token is considered valid if it matches the current counter or any
	// within the leeway.
	isValid := false
	for step := uint64(0); step <= h.Leeway; step++ {
		under := h.Counter - step

		expected, err := generateOTP(h.Key, under, h.Length, h.Algorithm)
		if err != nil {
			return false, err
		}
		if expected == token {
			isValid = true
			break
		}

		over := h.Counter + step
		expected, err = generateOTP(h.Key, over, h.Length, h.Algorithm)
		if err != nil {
			return false, err
		}
		if expected == token {
			isValid = true
			break
		}
	}

	if isValid {
		h.Counter++
	}

	return isValid, nil
}

// ensureDefaults applies sensible default values, if any of them is empty, so
// that the OTP generation works properly.
// Defaults:
//     - Algorithm = SHA1
//     - Length = 6
func (h *HOTP) ensureDefaults() {
	if h.Leeway == 0 {
		h.Leeway = HOTPDefaultLeeway
	}

	if h.Algorithm == 0 {
		h.Algorithm = HmacSHA1
	}

	if h.Length == 0 {
		h.Length = Length6
	}
}

// ensureKey generates a proper random key if no value is provided by the caller.
func (h *HOTP) ensureKey() (err error) {
	if h.Key != "" {
		return nil
	}

	h.Key, err = randomKey(RandomKeyLength)

	return err
}
