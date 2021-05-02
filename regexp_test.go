package main

import (
	"testing"
)

func TestSimpleEvaluation(t *testing.T) {
	r := new(regexp)
	err := r.addState(
		regexpopts{
			exp:        "f",
			recurrency: 1,
		},
		regexpopts{
			exp:        "o",
			recurrency: 1,
		},
		regexpopts{
			exp:        "r",
			recurrency: 1,
		},
	)

	if err != nil {
		t.Errorf("unexpected error %v\n", err)
	}

	tests := []struct {
		input  string
		expect bool
	}{
		{
			input:  "for",
			expect: true,
		},
		{
			input:  "while",
			expect: false,
		},
	}

	for _, test := range tests {
		eval := r.evaluate(test.input)
		if eval != test.expect {
			t.Errorf("Evaluate %s \t expected: %v \t got: %v\n", test.input, test.expect, eval)
		}
	}
}

func TestRangeEvaluation(t *testing.T) {
	r := new(regexp)

	err := r.addState(regexpopts{
		exp:        "a-c",
		recurrency: 3,
	}, regexpopts{
		exp:        "o",
		recurrency: 1,
	}, regexpopts{
		exp:        "r",
		recurrency: 3,
	})

	if err != nil {
		t.Errorf("unexpected error %v\n", err)
	}

	tests := []struct {
		input  string
		expect bool
	}{
		{
			input:  "abcorrr",
			expect: true,
		},
		{
			input:  "bcaorrr",
			expect: true,
		},
		{
			input:  "ior",
			expect: false,
		},
	}

	for _, test := range tests {
		eval := r.evaluate(test.input)
		if eval != test.expect {
			t.Errorf("Evaluate %s \t expected: %v \t got: %v\n", test.input, test.expect, eval)
		}
	}
}
