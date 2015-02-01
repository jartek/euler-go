package main

import "testing"

type testpair struct {
	input  int
	output int
}

var tests = []testpair{
	{3, 6},
	{5, 28},
	{500, 76576500},
}

func TestFirstTriangleNumber(t *testing.T) {
	for _, pair := range tests {
		num := FirstTriangleNumber(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
