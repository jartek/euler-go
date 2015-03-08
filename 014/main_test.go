package main

import "testing"

type testpair struct {
	input  int
	output int
}

var tests = []testpair{
	{10, 9},
	{5, 3},
	{1000000, 837799},
}

func TestLongestCollatzSequence(t *testing.T) {
	for _, pair := range tests {
		num := LongestCollatzSequence(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
