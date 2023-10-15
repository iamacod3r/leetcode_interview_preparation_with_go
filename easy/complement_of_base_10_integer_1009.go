package easy

import "interview_go/common"

type ComplementOfBase10Integer struct {
}

func (t *ComplementOfBase10Integer) bitwiseComplement(n int) int {

	if n == 0 {
		return 1
	}

	bitmask := n
	intRepeat := 1
	for intRepeat < 32 {
		bitmask |= (bitmask >> intRepeat)
		intRepeat *= 2
	}

	return bitmask ^ n
}

func (t *ComplementOfBase10Integer) bitwiseComplement2(n int) int {

	if n == 0 {
		return 1
	}

	num := n
	bit := 1

	for num > 0 {
		n = n ^ bit
		bit = bit << 1
		num = num >> 1
	}
	return n
}

func (t *ComplementOfBase10Integer) Test() {
	n1 := 5
	e1 := 2
	r1 := t.bitwiseComplement2(n1)
	common.PrintInt(r1, e1)

	n2 := 7
	e2 := 0
	r2 := t.bitwiseComplement2(n2)
	common.PrintInt(r2, e2)

	n3 := 10
	e3 := 5
	r3 := t.bitwiseComplement2(n3)
	common.PrintInt(r3, e3)
}
