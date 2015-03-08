package main

import "testing"

type testpair struct {
	input  float64
	output float64
}

var tests = []testpair{
	{2.0, 6.0},
	{20.0, 137846528820.0},
}

func TestLatticePaths(t *testing.T) {
	for _, pair := range tests {
		num := LatticePaths(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
