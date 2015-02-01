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

func LargestPrime(input float64) float64 {
	var i float64
	for i = math.Floor(math.Sqrt(input)); i > 1; i-- {
		if math.Mod(input, i) == 0 && isPrime(i) {
			return i
		}
	}
	return -1
}
