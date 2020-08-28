package main

import (
	"testing"
)

func TestHOTPExample(t *testing.T) {
	defer func() {
		if recover() != nil {
			t.Errorf("unexpected panic, package ./examples/hotp is expected to execute successfully")
		}
	}()

	main()
}
