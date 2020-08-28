package main

import (
	"fmt"
	"log"

	"github.com/jltorresm/otpgo"
)

func main() {
	fmt.Println("Time-Based One-Time Password")

	// Will use the default parameters and will generate a random key
	h := otpgo.TOTP{}

	// Generate standalone code
	otp1, err := h.Generate()
	if err != nil {
		log.Panicf("unexpected error when generating OTP: %s", err)
	}

	msg := "Used key: %s\nGenerated codes:\n\t- %s\n"
	fmt.Printf(msg, h.Key, otp1)

	// Validate a couple codes
	otherCode := "966205"

	ok, err := h.Validate(otp1)
	if err != nil {
		log.Panicf("unexpected error when validating OTP: %s", err)
	}

	ok2, err := h.Validate(otherCode)
	if err != nil {
		log.Panicf("unexpected error when validating OTP: %s", err)
	}

	fmt.Printf("Validated codes:\n\t- %s -> %v\n\t- %s -> %v\n", otp1, ok, otherCode, ok2)
}
