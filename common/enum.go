package common

import (
	"fmt"
)

type PersonEnum string

const (
	Iamacod3r PersonEnum = "Iamacod3r"
)

func PrintPersonEnum(d PersonEnum) {
	fmt.Printf("%s devloper\n", d)
}

func Test() {
	var who PersonEnum = Iamacod3r
	PrintPersonEnum(who)
	PrintPersonEnum("Iamacod3r")
}
