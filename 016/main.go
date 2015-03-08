package main

import "math/big"

func PowerDigitSum(num int64) int64 {
	var sum int64
	zero := big.NewInt(0)
	two := big.NewInt(2)
	ten := big.NewInt(10)

	mod := big.NewInt(0)
	n := big.NewInt(0)

	number := big.NewInt(num)

	n.Exp(two, number, nil)

	for zero.Cmp(n) < 0 {
		mod.Mod(n, ten)
		n.Quo(n, ten)
		sum = sum + mod.Int64()
	}

	return sum
}
