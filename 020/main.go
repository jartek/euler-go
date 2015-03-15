package main

import "math/big"

func FactorialSum(n int64) int64 {
	var sum int64
	zero := big.NewInt(0)
	ten := big.NewInt(10)
	mod := big.NewInt(0)

	factorial := big.NewInt(1)
	factorial.MulRange(int64(1), int64(n))

	for zero.Cmp(factorial) < 0 {
		mod.Mod(factorial, ten)
		factorial.Quo(factorial, ten)
		sum = sum + mod.Int64()
	}
	return sum
}
