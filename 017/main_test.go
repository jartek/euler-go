package main

import "testing"

type testpair struct {
	input  int
	output interface{}
}

func TestNumberLetterCounts(t *testing.T) {
	var tests = []testpair{
		{5, 19},
		{1000, 21124},
	}
	for _, pair := range tests {
		num := NumberLetterCounts(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}

func TestNumberInLetters(t *testing.T) {
	var tests = []testpair{
		{1, "one"},
		{11, "eleven"},
		{21, "twentyone"},
		{37, "thirtyseven"},
		{137, "onehundredandthirtyseven"},
		{342, "threehundredandfortytwo"},
		{959, "ninehundredandfiftynine"},
		{1000, "onethousand"},
	}
	for _, pair := range tests {
		num := NumberInLetters(pair.input)
		if num != pair.output {
			t.Error(
				"For", pair.input, "Expected", pair.output, "Got", num,
			)
		}
	}
}
