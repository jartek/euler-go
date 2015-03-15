package main

func TriangleSum(input [][]int) int {
	var max int
	for i := len(input) - 2; i >= 0; i-- {
		for j := range input[i] {
			max = input[i+1][j]
			if max < input[i+1][j+1] {
				max = input[i+1][j+1]
			}
			input[i][j] += max
		}
	}
	return input[0][0]
}
