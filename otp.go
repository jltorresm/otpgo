package otpgo

import (
	"crypto/hmac"
	"crypto/rand"
	"encoding/base32"
	"encoding/binary"
)

// RandomKeyLength is the recommended length for the key used to generate OTPs.
// This length will be used to generate default keys (in HOTP.EnsureKey and
// TOTP.EnsureKey), when the caller does not provide one explicitly.
const RandomKeyLength = 64

// Generates a new OTP using the specified parameters based on the rfc4226.
func generateOTP(key string, counter uint64, length otpLength, algorithm hmacAlgorithm) (string, error) {
	// Decode secret key to bytes
	k, err := base32.StdEncoding.DecodeString(key)
	if err != nil {
		return "", ErrorInvalidKey{msg: err.Error()}
	}

	// Convert the counter to bytes
	msg := make([]byte, 8)
	binary.BigEndian.PutUint64(msg, counter)

	// Start the hmac algorithm
	hm := hmac.New(algorithm.Hash, k)
	hm.Write(msg)
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
func randomKey(length int) (string, error) {
	buff := make([]byte, length)
	if _, err := rand.Read(buff); err != nil {
		return "", err
	}

	return base32.StdEncoding.EncodeToString(buff), nil
}
