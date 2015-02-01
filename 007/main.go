package main

import "math"

func isPrime(num float64) bool {
	var i float64

	for i = math.Floor(math.Sqrt(num)); i > 1; i-- {
		if math.Mod(num, i) == 0 {
			return false
		}
	}

	return true
}

func NthPrime(n float64) float64 {
	var i, num float64
	num = 1
	for n != 0 {
		for i = num + 1; ; i++ {
			if (i == 2 || math.Mod(i, 2) != 0) && isPrime(i) {
				n = n - 1
				break
			}
		}
		num = i
	}
	return num
}
