package config

import (
	"fmt"
	"math"
	"strconv"
)

// Length type describes the recommended OTP lengths.
type Length int

// Supported length to the OTP generated using HOTP or TOTP.
const (
	Length1 Length = iota + 1
	Length2
	Length3
	Length4
	Length5
	Length6
	Length7
	Length8
)

// Truncate will cut the provided number to fit the Length.
func (l Length) Truncate(number int) int {
	return number % int(math.Pow10(int(l)))
}

// LeftPad adds extra zeroes to the left of the number to complete the Length.
func (l Length) LeftPad(number int) string {
	format := "%0" + strconv.Itoa(int(l)) + "d"
	return fmt.Sprintf(format, number)
}

// String converts the Length into a string.
func (l Length) String() string {
	return strconv.Itoa(int(l))
}
