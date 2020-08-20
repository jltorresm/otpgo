package otpgo

import (
	"fmt"
)

// The ErrorInvalidKey represents an invalid key used to generate OTPs.
type ErrorInvalidKey struct {
	msg string
}

func (eik ErrorInvalidKey) Error() string {
	return fmt.Sprintf("invalid key: %s", eik.msg)
}
