package main

import "testing"

func TestCorrectParentheses(t *testing.T) {
	for _, tc := range []struct {
		s    string
		n    int
		want bool
	}{
		{"{}", 2, true},
		{"()(}", 4, false},
		{"([{}]{]})", 9, false},
		{"]", 1, false},
		{"({{{{(", 6, false},
		{"(){}[]", 6, true},
		{"([{}])", 6, true},
		{"[[]", 3, false},
		{"()([])))", 9, false},
		{"(){[]}", 6, true},
	} {
		if res := correctParentheses(tc.s, tc.n); res != tc.want {
			t.Errorf("got =%v, want = %v", res, tc.want)
		}
	}
}
