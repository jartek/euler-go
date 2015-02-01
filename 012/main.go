package main

import "math"

func NumOfDivisors(n int) int {
	num := 0
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			num = num + 2
		}
	}
	return num
}

func FirstTriangleNumber(limit int) int {
	n := 1
	triangularNumber := 0
	numOfDivisors := 0
	for numOfDivisors < limit {
		triangularNumber = n * (n + 1) / 2
		numOfDivisors = NumOfDivisors(triangularNumber)
		n = n + 1
	}
	return triangularNumber
}
