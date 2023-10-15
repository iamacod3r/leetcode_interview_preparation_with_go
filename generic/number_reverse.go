package generic

import (
	"math"
)

func ReverseNumber(x int) int {
	sign := 1

	if x < 0 {
		sign = -1
	}

	x = int(math.Abs(float64(x)))

	var reverseDigit int

	for x > 0 {
		lastDigit := x % 10
		reverseDigit = reverseDigit*10 + lastDigit
		x = x / 10
	}

	if sign == -1 {
		reverseDigit = reverseDigit * -1
	}

	return reverseDigit
}
