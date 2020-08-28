package main

import (
	"testing"
)

func TestTOTPExample(t *testing.T) {
	defer func() {
		if recover() != nil {
			t.Errorf("unexpected panic, package ./examples/totp is expected to execute successfully")
		}
	}()

	main()
}
