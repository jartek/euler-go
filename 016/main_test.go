package main

import "testing"

type testpair struct {
	input  int64
	output int64
}

var tests = []testpair{
	{15, 26},
	{1000, 1366},
}

func TestPowerDigitSum(t *testing.T) {
	for _, pair := range tests {
		num := PowerDigitSum(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
