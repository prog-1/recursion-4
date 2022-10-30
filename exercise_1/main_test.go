package main

import (
	"fmt"
	"testing"
)

func TestCorrectParentheses(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  bool
	}{
		{")(", false},
		{"(()())(", false},
		{")()(", false},
		{"({[}])", false},
		{"(}[])", false},
		{"()", true},
		{"(()())()", true},
		{"()()()", true},
		{"()[]{}", true},
		{"([{}])", true},
		{"(([{}]))", true},
	} {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			if got := CorrectParentheses(tc.input); got != tc.want {
				t.Errorf("CorrectParentheses( %v ) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
