package main

import "testing"

type testpair struct {
	input  float64
	output float64
}

var tests = []testpair{
	{10, 2640},
	{100, 25164150},
}

func TestSmallestMultiple(t *testing.T) {
	for _, pair := range tests {
		num := SumSquareDifference(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
