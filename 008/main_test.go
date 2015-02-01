package main

import "testing"

type testpair struct {
	input  int
	output int
}

var tests = []testpair{
	{4, 5832},
	{13, 23514624000},
}

func TestGreatestNDigit(t *testing.T) {
	for _, pair := range tests {
		num := GreatestNDigit(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
