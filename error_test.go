package otpgo

import (
	"testing"
)

func TestErrorInvalidKey_Error(t *testing.T) {
	err := ErrorInvalidKey{msg: "an arbitrary error message"}
	expectedError := "invalid key: an arbitrary error message"

	if err.Error() != expectedError {
		t.Errorf("unexpected error\nexpected: %s\n  actual: %s", expectedError, err.Error())
	}
}
