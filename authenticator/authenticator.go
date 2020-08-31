package authenticator

import (
	"fmt"
	"net/url"
)

// The KeyUri type holds all the config necessary to generate a valid
// registration URI for any authenticator app, e.g.: Google Authenticator, Authy,
// etc.
type KeyUri struct {
	Type       string // hotp | totp
	Label      Label
	Parameters paramsFormatter
}

// String encodes all the KeyUri info as a valid URI.
func (ku *KeyUri) String() string {
	format := "otpauth://%s/%s?%s"
	raw := fmt.Sprintf(format, ku.Type, ku.Label.String(), ku.Parameters.AsUrlValues(ku.Label.Issuer).Encode())
	uri, _ := url.Parse(raw)
	return uri.String()
}

// The Label is used to identify which account a key is associated with.
type Label struct {
	AccountName string // URI encoded
	Issuer      string
}

// The String method will encode the Label in a valid format to be included in
// the key URI.
func (l *Label) String() string {
	return fmt.Sprintf("%s:%s", l.Issuer, l.AccountName)
}

// The paramsFormatter type has the ability to encode the OTP params into
// expected formats to be used for exporting.
type paramsFormatter interface {
	AsUrlValues(issuer string) url.Values
}
