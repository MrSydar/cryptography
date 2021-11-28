package bbs

import (
	"errors"
)

func areCongruent(a int, b int, n int) bool {
	return (a-b)%n == 0
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func checkInput(p int, q int, x int) error {
	if !areCongruent(p, 3, 4) {
		return errors.New("p and 3 mod 4 are not congruent")
	} else if !areCongruent(q, 3, 4) {
		return errors.New("q and 3 mod 4 are not congruent")
	} else if gcd(p*q, x) != 1 {
		return errors.New("p * q product and x are not comprime")
	} else {
		return nil
	}
}

func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

func BBS(p int, q int, x int, n int) (bits []bool, err error) {
	if err := checkInput(p, q, x); err != nil {
		return nil, err
	}

	bits = make([]bool, n)

	N := p * q

	x = mod(x*x, N)
	bits[0] = (x & 1) == 1
	for i := 1; i < n; i++ {
		x = mod(x*x, N)
		bits[i] = (x & 1) == 1
	}

	return bits, nil
}
