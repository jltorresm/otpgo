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
	otp35, err := h.Generate()
	if err != nil {
		log.Panicf("unexpected error when generating OTP: %s", err)
	}

	otp36, err := h.Generate()
	if err != nil {
		log.Panicf("unexpected error when generating OTP: %s", err)
	}

	otp37, err := h.Generate()
	if err != nil {
		log.Panicf("unexpected error when generating OTP: %s", err)
	}

	msg := "Used key: %s\nGenerated codes:\n\t- %s\n\t- %s\n\t- %s\n"
	fmt.Printf(msg, h.Key, otp35, otp36, otp37)

	// Validate the last code
	//     Since the token generation should be done in an external device, we
	//     are rewinding the Counter one time.
	h.Counter--
	isValid1, err := h.Validate(otp37)
	if err != nil {
		log.Panicf("unexpected error when validating OTP: %s", err)
	}

	invalidToken := "123456"
	isValid2, err := h.Validate(invalidToken)
	if err != nil {
		log.Panicf("unexpected error when validating OTP: %s", err)
	}

	msg = "Validated code:\n\t- %s --> %v\n\t- %s --> %v\n"
	fmt.Printf(msg, otp37, isValid1, invalidToken, isValid2)
}
