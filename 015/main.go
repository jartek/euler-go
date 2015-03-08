package main

func factorial(num float64) float64 {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

func LatticePaths(num float64) float64 {
	m := factorial(2 * num)
	n := factorial(num)
	return m / (n * n)
}
