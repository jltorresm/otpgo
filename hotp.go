package otpgo

type HOTP struct {
	Key       string // Secret
	Counter   uint64
	Algorithm hmacAlgorithm
	Length    otpLength
}

// Generates the HMAC-Based One-Time Password.
func (h *HOTP) Generate() (string, error) {
	// Make sure we have sensible values to generate secure OTPs
	h.EnsureDefaults()

	// Make sure we have a valid non-empty key
	if err := h.EnsureKey(); err != nil {
		return "", err
	}

	return generateOTP(h.Key, h.Counter, h.Length, h.Algorithm)
}

// If any required value is empty, it will apply sensible default values so that
// the OTP generation works properly.
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

// If no key is provided by the caller, it generates a valid random key.
func (h *HOTP) EnsureKey() (err error) {
	if h.Key != "" {
		return nil
	}

	h.Key, err = randomKey(RandomKeyLength)

	return err
}
