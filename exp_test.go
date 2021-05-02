package main

import "testing"

func TestRangedExpression(t *testing.T) {
	tests := []struct {
		e            exp
		hasRange     bool
		isValidRange bool
		getRange     string
	}{
		{
			e:            "a",
			hasRange:     false,
			isValidRange: false,
			getRange:     "",
		},
		{
			e:            "a-z",
			hasRange:     true,
			isValidRange: true,
			getRange:     alphabet[0:26],
		},
		{
			e:            "z-a",
			hasRange:     true,
			isValidRange: false,
			getRange:     "",
		},
	}

	for _, test := range tests {
		hr := test.e.hasRange()
		valid := test.e.isValidRange()
		getrange := test.e.getRange()

		if hr != test.hasRange {
			t.Errorf("Evaluate %s.hasRange() \t expected: %v \t got: %v\n", test.e, test.hasRange, hr)
		}

		if valid != test.isValidRange {
			t.Errorf("Evaluate %s.isValidRange() \t expected: %v \t got: %v\n", test.e, test.isValidRange, valid)
		}

		if getrange != test.getRange {
			t.Errorf("Evaluate %s.getRange() \t expected: %v \t got: %v\n", test.e, test.getRange, getrange)
		}
	}
}
