package main

import "testing"

type testpair struct {
	input  float64
	output float64
}

var tests = []testpair{
	{10, 17},
	{2000000, 142913828922},
}

func TestSumOfPrimes(t *testing.T) {
	for _, pair := range tests {
		num := SumOfPrimes(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
