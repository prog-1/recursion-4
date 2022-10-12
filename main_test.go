package main

import (
	"testing"
)

func TestCorrectParentheses(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want bool
	}{
		{"(()())(", false},
		{")()(", false},
		{"(()())()", true},
		{"()()()", true},
		{"({[}])", false},
		{"(}[])", false},
		{"()[]{}", true},
		{"([{}])", true},
		{"{}(({}[}}}}]}]{{)]{)", false},
		{"{[()]}", true},
	} {
		got := CorrectParentheses(tc.s)
		if got != tc.want {
			t.Errorf("CorrectParentheses(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}
