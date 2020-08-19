package otpgo

import (
	"fmt"
	"math"
	"strconv"
)

type otpLength int

const (
	Length1 otpLength = iota + 1
	Length2
	Length3
	Length4
	Length5
	Length6
	Length7
	Length8
)

func (l otpLength) Truncate(number int) int {
	return number % int(math.Pow10(int(l)))
}

func (l otpLength) LeftPad(number int) string {
	format := "%0" + strconv.Itoa(int(l)) + "d"
	return fmt.Sprintf(format, number)
}
