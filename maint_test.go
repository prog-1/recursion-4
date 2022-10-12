package main

import "testing"

func TestCorrect(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  bool
	}{
		{"()", true},
		{"(){)}", false},
		{"()(}", false},
		{")", false},
		{"(", false},
		{"({)}", false},
		{"(){}[]", true},
		{"[({})]", true},
		{"()(", false},
		{"{{()}", false},
		{"(){[]}", true},
	} {
		if CorrectParentheses(tc.input) != tc.want {
			t.Errorf("CorrectParentheses(%v) = %v, want = %v", tc.input, CorrectParentheses(tc.input), tc.want)
		}
	}
}
