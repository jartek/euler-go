package main

import "testing"

type testpair struct {
	input  int64
	output int64
}

func TestFactorialSum(t *testing.T) {
	var tests = []testpair{
		{3, 6},
		{10, 27},
		{100, 648},
	}

	for _, pair := range tests {
		num := FactorialSum(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
