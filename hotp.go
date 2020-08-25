package otpgo

// The HOTP type used to generate HMAC-Based One-Time Passwords.
type HOTP struct {
	Key       string // Secret
	Counter   uint64
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

	h.Counter++

	return generateOTP(h.Key, h.Counter, h.Length, h.Algorithm)
}

// ensureDefaults applies sensible default values, if any of them is empty, so
// that the OTP generation works properly.
// Defaults:
//     - Algorithm = SHA1
//     - Length = 6
func (h *HOTP) ensureDefaults() {
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
