package main

import "testing"

func TestIsLeapYear(t *testing.T) {
	type testpair struct {
		input  int
		output bool
	}

	var tests = []testpair{
		{1900, false},
		{1904, true},
		{2000, true},
	}

	for _, pair := range tests {
		num := isLeapYear(pair.input, 1)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}

func TestFirstDayOfTheYear(t *testing.T) {
	type testpair struct {
		input  int
		output int
	}

	var tests = []testpair{
		{1900, 1},
		{1901, 2},
		{1902, 3},
		{1903, 4},
		{1904, 5},
		{1905, 0},
		{2000, 6},
	}
	for _, pair := range tests {
		num := FirstDayOfTheYear(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}

func TestNumberOfSundays(t *testing.T) {
	type testpair struct {
		input  []int
		output int
	}

	var tests = []testpair{
		{[]int{1901, 2000}, 171},
	}
	for _, pair := range tests {
		num := NumberOfSundays(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
