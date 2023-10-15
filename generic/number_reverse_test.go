package generic

import (
	"log"
	"testing"
)

func TestReverseNumber(t *testing.T) {
	reversedNumber := ReverseNumber(123)
	log.Println(reversedNumber)

	reversedNumber = ReverseNumber(140)
	log.Println(reversedNumber)

	reversedNumber = ReverseNumber(-123)
	log.Println(reversedNumber)

	reversedNumber = ReverseNumber(0)
	log.Println(reversedNumber)
}
