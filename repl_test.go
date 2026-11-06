package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		}, // add more cases here
		{
			input:    "hello, world",
			expected: []string{"hello,", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "word",
			expected: []string{"word"},
		},
		{
			input:    "3 words here",
			expected: []string{"3", "words", "here"},
		},
		{
			input:    "UPPER CASE TO LOWER",
			expected: []string{"upper", "case", "to", "lower"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			// error and continue here
			t.Errorf("Length did not match; actual: %v; expected: %v", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Word %s did not match %s; Expected: %v, Actual %v", word, expectedWord, c.expected, actual)
			}
		}
	}
}
