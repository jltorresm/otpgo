package main

import (
	"fmt"
	"log"

	"github.com/jltorresm/otpgo"
)

func main() {
	fmt.Println("HMAC-Based One-Time Password")

	// Will use the default parameters and will generate a random key
	h := otpgo.HOTP{Counter: 34}

	// Generate standalone codes
	otp35, _ := h.Generate()

	h.Counter++ // Counter has to manually increased if you are not calling Validate()
	otp36, _ := h.Generate()

	// Validate the last code
	isValid1, err := h.Validate(otp36)
	if err != nil {
		log.Panicf("unexpected error when validating OTP: %s", err)
	}

	// This code will be the next one, even though we didn't call h.Counter++
	// because successful validation will internally make the counter to go to
	// the next expected counter.
	otp37, _ := h.Generate()

	msg := "Used key: %s\nGenerated codes:\n\t- %s\n\t- %s\n\t- %s\n"
	fmt.Printf(msg, h.Key, otp35, otp36, otp37)

	invalidToken := "123456"
	isValid2, err := h.Validate(invalidToken)
	if err != nil {
		log.Panicf("unexpected error when validating OTP: %s", err)
	}

	msg = "Validated code:\n\t- %s --> %v\n\t- %s --> %v\n"
	fmt.Printf(msg, otp36, isValid1, invalidToken, isValid2)

	// If trying to validate without a key it will error out.
	h = otpgo.HOTP{}
	isValid, err := h.Validate("a-token")
	fmt.Printf("Trying to validate without key, is valid: %v, error: %s\n", isValid, err)
}
