package main

import "testing"

type testpair struct {
	input  int
	output int
}

var tests = []testpair{
	{10, 23},
	{1000, 233168},
}

func TestMultiples(t *testing.T) {
	for _, pair := range tests {
		sum := Multiples(pair.input)
		if sum != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", sum,
			)
		}
	}
}
