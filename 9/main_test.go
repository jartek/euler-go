package main

import "testing"

type testpair struct {
	input  float64
	output int64
}

var tests = []testpair{
	{1000, 31875000},
}

func TestPythagoreanTriplet(t *testing.T) {
	for _, pair := range tests {
		num := PythagoreanTriplet(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
