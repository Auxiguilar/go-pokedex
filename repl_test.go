package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for i, c := range cases {
		want := c.expected
		got := cleanInput(c.input)

		if len(want) != len(got) {
			t.Errorf("fail test %d: length: %d != %d\nwanted %v, got %v", i+1, len(want), len(got), want, got)
		}

		for i := range got {
			if i > len(want)-1 {
				break
			}

			gotWord := got[i]
			wantWord := c.expected[i]

			if gotWord != wantWord {
				t.Errorf("fail: wanted %v, got %v", want, got)
			}
		}
	}
}
