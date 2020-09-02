package otpgo

import (
	"crypto/hmac"
	"crypto/rand"
	"encoding/base32"
	"encoding/binary"
	"strings"

	"github.com/jltorresm/otpgo/config"
)

// RandomKeyLength is the recommended length for the key used to generate OTPs.
// This length will be used to generate default keys (in HOTP.ensureKey and
// TOTP.ensureKey), when the caller does not provide one explicitly.
const RandomKeyLength = 64

var otpBase32Encoding = base32.StdEncoding.WithPadding(base32.NoPadding)

// Generates a new OTP using the specified parameters based on the rfc4226.
func generateOTP(key string, counter uint64, length config.Length, algorithm config.HmacAlgorithm) (string, error) {
	// Ensure key is uppercase
	key = strings.ToUpper(key)

	// Trim unnecessary paddings in case the key was generated externally.
	key = strings.TrimRight(key, string(base32.StdPadding))

	// Decode secret key to bytes
	k, err := otpBase32Encoding.DecodeString(key)
	if err != nil {
		return "", ErrorInvalidKey{msg: err.Error()}
	}

	// Convert the counter to bytes
	msg := make([]byte, 8)
	binary.BigEndian.PutUint64(msg, counter)

	// Start the hmac algorithm
	hm := hmac.New(algorithm.Hash, k)
	if _, err := hm.Write(msg); err != nil {
		return "", err
	}
	sum := hm.Sum([]byte{})

	// Build the result integer
	offset := sum[len(sum)-1] & 0xf

	bin := ((int(sum[offset]) & 0x7f) << 24) |
		((int(sum[offset+1]) & 0xff) << 16) |
		((int(sum[offset+2]) & 0xff) << 8) |
		(int(sum[offset+3]) & 0xff)

	rawOtp := length.Truncate(bin)
	otp := length.LeftPad(rawOtp)

	return otp, nil
}

// Generates a random key of the specified length, usable for OTP generation.
func randomKey(length uint) (string, error) {
	buff := make([]byte, length)
	if _, err := rand.Read(buff); err != nil {
		return "", err
	}

	return otpBase32Encoding.EncodeToString(buff), nil
}
