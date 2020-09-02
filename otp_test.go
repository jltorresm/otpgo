package otpgo

import (
	"testing"
)

func TestRandomKey(t *testing.T) {
	cases := []struct {
		label          string
		length         uint
		expectedLength int
		expectedError  error
	}{
		{"Zero", 0, 0, nil},
		{"One", 1, 2, nil},
		{"Three", 3, 5, nil},
		{"Small", 10, 16, nil},
		{"Normal", 64, 103, nil},
		{"Big", 1024, 1639, nil},
	}

	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			key, err := randomKey(c.length)

			if c.expectedLength != len(key) {
				t.Errorf("unexpected key length\nexpected: %d\n  actual: %d", c.expectedLength, len(key))
			}

			if c.expectedError != err {
				t.Errorf("unexpected error\nexpected: %s\n  actual: %s", c.expectedError, err)
			}

			if len(key) > 0 && key[len(key)-1:] == "=" {
				t.Errorf("unexpected padding, expected the key to have no padding")
			}
		})
	}
}
