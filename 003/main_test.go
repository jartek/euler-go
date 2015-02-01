package main

import "testing"

type testpair struct {
	input  float64
	output float64
}

var tests = []testpair{
	{45, 5},
	{13195, 29},
	{600851475143, 6857},
}

func TestLargestPrime(t *testing.T) {
	for _, pair := range tests {
		num := LargestPrime(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
