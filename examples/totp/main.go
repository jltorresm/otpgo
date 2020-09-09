package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jltorresm/otpgo"
)

func main() {
	fmt.Println("Time-Based One-Time Password")

	// Will use the default parameters and will generate a random key
	t := otpgo.TOTP{}

	// Generate standalone code
	otp1, _ := t.Generate()

	msg := "Used key: %s\nGenerated codes:\n\t- %s\n"
	fmt.Printf(msg, t.Key, otp1)

	// Validate a couple codes
	otherCode := "966205"

	ok, err := t.Validate(otp1)
	if err != nil {
		log.Panicf("unexpected error when validating OTP: %s", err)
	}

	ok2, err := t.Validate(otherCode)
	if err != nil {
		log.Panicf("unexpected error when validating OTP: %s", err)
	}

	fmt.Printf("Validated codes:\n\t- %s -> %v\n\t- %s -> %v\n", otp1, ok, otherCode, ok2)

	// If trying to validate without a key it will error out.
	t2 := otpgo.TOTP{}
	isValid, err := t2.Validate("a-token")
	fmt.Printf("Trying to validate without key, is valid: %v, error: %s\n", isValid, err)

	// To export the secrets you will need to the KeyUri based on your TOTP variable.
	// For this purpose an account name and an issuer are needed.
	aUsername := "username@example.com"
	anIssuer := "A Company"
	ku := t.KeyUri(aUsername, anIssuer)

	// From here you can get the plain text uri.
	msg = "Exporting config for \"%s\" at \"%s\":\n" +
		"\t- Plain URI --> %s\n" +
		"\t- QR code image, base 64 encoded (truncated to save space) --> %s...\n" +
		"\t- JSON --> %s\n"
	qr, _ := ku.QRCode()
	jsonParams, _ := json.Marshal(ku)
	fmt.Printf(msg, aUsername, anIssuer, ku.String(), qr[0:200], jsonParams)
}
