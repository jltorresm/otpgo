package otpgo

import (
	"errors"
	"math"
	"net/url"
	"strconv"
	"time"

	"github.com/jltorresm/otpgo/authenticator"
	"github.com/jltorresm/otpgo/config"
)

const (
	// TOTPDefaultPeriod is the default time period to use if none is provided by the caller.
	TOTPDefaultPeriod = 30
	// TOTPDefaultDelay is the default acceptable delay window. A value of 1
	// means the OTP will be valid if it coincides with the calculated token for
	// the current time step, the next one or the one before.
	TOTPDefaultDelay = 1
)

// The TOTP type used to generate Time-Based One-Time Passwords.
type TOTP struct {
	Key       string // Secret base32 encoded string
	Period    int    // Number of seconds the TOTP is valid
	Delay     int    // Acceptable steps for network delay
	Algorithm config.HmacAlgorithm
	Length    config.Length
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
	counter := t.getCounter(time.Now().Unix())

	return generateOTP(t.Key, counter, t.Length, t.Algorithm)
}

// Validate will try to check if the provided token is a valid OTP for the
// current TOTP config.
//
// If the TOTP struct is using all the default values the config will be
// compatible with the Google Authenticator app, as well as most other apps.
func (t *TOTP) Validate(token string) (bool, error) {
	// This will be the base for all validations
	now := time.Now().Unix()

	// Validating without a proper key shouldn't happen
	if t.Key == "" {
		return false, errors.New("missing secret key for validation")
	}

	// Make sure we have sensible values to generate secure OTPs
	t.ensureDefaults()

	// Now go through all the possible valid tokens
	for step := 0; step <= t.Delay; step++ {
		pad := int64(t.Period * step)
		under := t.getCounter(now - pad)

		expected, err := generateOTP(t.Key, under, t.Length, t.Algorithm)
		if err != nil {
			return false, err
		}
		if expected == token {
			return true, nil
		}

		over := t.getCounter(now + pad)
		expected, err = generateOTP(t.Key, over, t.Length, t.Algorithm)
		if err != nil {
			return false, err
		}
		if expected == token {
			return true, nil
		}
	}

	return false, nil
}

// KeyUri return an authenticator.KeyUri configured with the current TOTP params.
//     - accountName is the username or email of the account
//     - issuer is the site or org
func (t *TOTP) KeyUri(accountName, issuer string) *authenticator.KeyUri {
	return &authenticator.KeyUri{
		Type: "totp",
		Label: authenticator.Label{
			AccountName: accountName,
			Issuer:      issuer,
		},
		Parameters: t,
	}
}

// AsUrlValues returns the TOTP parameters represented as url.Values.
func (t *TOTP) AsUrlValues(issuer string) url.Values {
	params := url.Values{}
	params.Add("secret", t.Key)
	params.Add("period", strconv.Itoa(t.Period))
	params.Add("algorithm", t.Algorithm.String())
	params.Add("digits", t.Length.String())
	params.Add("issuer", issuer)

	return params
}

// ensureDefaults applies sensible default values, if any of them is empty, so
// that the OTP generation works properly.
// Defaults:
//     - Period = TOTPDefaultPeriod = 30
//     - Delay = TOTPDefaultDelay = 1
//     - Algorithm = SHA1
//     - Length = 6
func (t *TOTP) ensureDefaults() {
	if t.Period == 0 {
		t.Period = TOTPDefaultPeriod
	}

	if t.Delay == 0 {
		t.Delay = TOTPDefaultDelay
	}

	if t.Algorithm == 0 {
		t.Algorithm = config.HmacSHA1
	}

	if t.Length == 0 {
		t.Length = config.Length6
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

// getCounter returns a valid counter based on the given timestamp.
func (t *TOTP) getCounter(timestamp int64) uint64 {
	return uint64(math.Floor(float64(timestamp) / float64(t.Period)))
}
