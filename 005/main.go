package main

import "math"

func SmallestMultiple(limit float64) float64 {
	var num, i float64 = 1, 2
	for ; i < limit; i++ {
		num = lcm(num, i)
	}
	return num
}

func lcm(a float64, b float64) float64 {
	return a * b / gcd(a, b)
}

func gcd(a float64, b float64) float64 {
	c := math.Floor(math.Mod(a, b))

	if c == 0 {
		return b
	}

	return gcd(b, c)
}
