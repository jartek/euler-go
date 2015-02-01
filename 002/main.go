package main

func FibonacciSum(limit int) int {
	var sum, tmp int
	series := []int{1, 2}

	for series[1] < limit {
		if series[1]%2 == 0 {
			sum = sum + series[1]
		}
		tmp = series[0]
		series[0] = series[1]
		series[1] = tmp + series[1]
	}

	return sum
}
