package main

import "testing"

type testpair struct {
	input  int
	output int
}

var tests = []testpair{
	{10, 10},
	{40, 44},
	{4000000, 4613732},
}

func TestFibonacciSum(t *testing.T) {
	for _, pair := range tests {
		sum := FibonacciSum(pair.input)
		if sum != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", sum,
			)
		}
	}
}
