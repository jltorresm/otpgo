package authenticator

import (
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/skip2/go-qrcode"
)

const (
	// QRSize indicates the size of the generated image containing the QR code.
	QRSize = 256
)

// The KeyUri type holds all the config necessary to generate a valid
// registration URI for any authenticator app, e.g.: Google Authenticator, Authy,
// etc.
type KeyUri struct {
	Type       string          `json:"type"` // hotp | totp
	Label      Label           `json:"label"`
	Parameters paramsFormatter `json:"parameters"`
}

// String encodes all the KeyUri info as a valid URI.
func (ku *KeyUri) String() string {
	format := "otpauth://%s/%s?%s"
	raw := fmt.Sprintf(format, ku.Type, ku.Label.String(), ku.Parameters.AsUrlValues(ku.Label.Issuer).Encode())
	uri, _ := url.Parse(raw)
	return uri.String()
}

// QRCode will encode the value returned by KeyUri.String into a base64
// encoded image containing a QR code that can be displayed and then scanned by
// the user. The return value is the base64 encoded image data.
func (ku *KeyUri) QRCode() (string, error) {
	uri := ku.String()

	qr, err := qrcode.New(uri, qrcode.Medium)
	if err != nil {
		return "", err
	}

	bytes, err := qr.PNG(QRSize)
	if err != nil {
		return "", err
	}

	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(bytes), nil
}

// The Label is used to identify which account a key is associated with.
type Label struct {
	AccountName string `json:"accountName"` // Should be a username, email, etc.
	Issuer      string `json:"issuer"`      // Should be the domain, company, org that is issuing the auth
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
