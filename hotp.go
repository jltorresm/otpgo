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
	h.EnsureDefaults()

	// Make sure we have a valid non-empty key
	if err := h.EnsureKey(); err != nil {
		return "", err
	}

	return generateOTP(h.Key, h.Counter, h.Length, h.Algorithm)
}

// EnsureDefaults applies sensible default values, if any of them is empty, so
// that the OTP generation works properly.
// Defaults:
//     - Algorithm = SHA1
//     - Length = 6
func (h *HOTP) EnsureDefaults() {
	if h.Algorithm == 0 {
		h.Algorithm = HmacSHA1
	}

	if h.Length == 0 {
		h.Length = Length6
	}
}

// EnsureKey generates a proper random key if no value is provided by the caller.
func (h *HOTP) EnsureKey() (err error) {
	if h.Key != "" {
		return nil
	}

	h.Key, err = randomKey(RandomKeyLength)

	return err
}
