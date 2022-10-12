package main

import (
	"reflect"
	"testing"
)

func TestCheck(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  bool
	}{
		{"([{}])", true},
		{"({[}])", false},
	} {
		if got := CorrectParentheses(tc.input); got != tc.want {
			t.Error("Got:", got, "Want:", tc.want)
		}
	}
}

func TestGen(t *testing.T) {
	for _, tc := range []struct {
		input uint32
		want  []string
	}{
		{2, []string{"(())", "([])", "({})", "()()", "()[]", "(){}", "[()]", "[[]]", "[{}]", "[]()", "[][]", "[]{}", "{()}", "{[]}", "{{}}", "{}()", "{}[]", "{}{}"}},
	} {
		var got []string
		GenParentheses(int(tc.input), func(s string) { got = append(got, s) })
		if !reflect.DeepEqual(got, tc.want) {
			t.Error("\nGot:", got, "\nWant:", tc.want)
		}
	}
}
