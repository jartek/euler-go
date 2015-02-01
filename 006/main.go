package main

import "math"

func SumSquareDifference(limit float64) float64 {
	return squareOfSum(limit) - sumOfSquares(limit)
}

func sumOfSquares(limit float64) float64 {
	return limit * (limit + 1) * (2*limit + 1) / 6
}

func squareOfSum(limit float64) float64 {
	return math.Pow((limit * (limit + 1) / 2), 2)
}
