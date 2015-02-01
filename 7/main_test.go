package main

import "testing"

type testpair struct {
	input  float64
	output float64
}

var tests = []testpair{
	{6, 13},
	{7, 17},
	{10001, 17},
}

func TestSmallestMultiple(t *testing.T) {
	for _, pair := range tests {
		num := NthPrime(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
