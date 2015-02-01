package main

import "testing"

type testpair struct {
	input  float64
	output float64
}

var tests = []testpair{
	{2, 9009},
	{3, 906609},
}

func TestLargestPrime(t *testing.T) {
	for _, pair := range tests {
		num := LargestPalindrome(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
