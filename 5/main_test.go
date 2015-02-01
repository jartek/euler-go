package main

import "testing"

type testpair struct {
	input  float64
	output float64
}

var tests = []testpair{
	{10, 2520},
	{20, 232792560},
}

func TestSmallestMultiple(t *testing.T) {
	for _, pair := range tests {
		num := SmallestMultiple(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
