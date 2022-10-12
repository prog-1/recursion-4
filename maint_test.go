package main

import (
	"reflect"
	"testing"
)

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

func TestGenerator(t *testing.T) {
	for _, tc := range []struct {
		input int
		want  []string
	}{
		{2, []string{"(())", "({})", "([])", "()()", "(){}", "()[]", "{()}", "{{}}", "{[]}", "{}()", "{}{}", "{}[]", "[()]", "[{}]", "[[]]", "[]()", "[]{}", "[][]"}},
		{3, []string{"((()))", "(({}))", "(([]))", "(()())", "((){})", "(()[])", "(())()", "(()){}", "(())[]", "({()})", "({{}})", "({[]})", "({}())", "({}{})", "({}[])", "({})()", "({}){}", "({})[]", "([()])", "([{}])", "([[]])", "([]())", "([]{})", "([][])", "([])()", "([]){}", "([])[]", "()(())", "()({})", "()([])", "()()()", "()(){}", "()()[]", "(){()}", "(){{}}", "(){[]}", "(){}()", "(){}{}", "(){}[]", "()[()]", "()[{}]", "()[[]]", "()[]()", "()[]{}", "()[][]", "{(())}", "{({})}", "{([])}", "{()()}", "{(){}}", "{()[]}", "{()}()", "{()}{}", "{()}[]", "{{()}}", "{{{}}}", "{{[]}}", "{{}()}", "{{}{}}", "{{}[]}", "{{}}()", "{{}}{}", "{{}}[]", "{[()]}", "{[{}]}", "{[[]]}", "{[]()}", "{[]{}}", "{[][]}", "{[]}()", "{[]}{}", "{[]}[]", "{}(())", "{}({})", "{}([])", "{}()()", "{}(){}", "{}()[]", "{}{()}", "{}{{}}", "{}{[]}", "{}{}()", "{}{}{}", "{}{}[]", "{}[()]", "{}[{}]", "{}[[]]", "{}[]()", "{}[]{}", "{}[][]", "[(())]", "[({})]", "[([])]", "[()()]", "[(){}]", "[()[]]", "[()]()", "[()]{}", "[()][]", "[{()}]", "[{{}}]", "[{[]}]", "[{}()]", "[{}{}]", "[{}[]]", "[{}]()", "[{}]{}", "[{}][]", "[[()]]", "[[{}]]", "[[[]]]", "[[]()]", "[[]{}]", "[[][]]", "[[]]()", "[[]]{}", "[[]][]", "[](())", "[]({})", "[]([])", "[]()()", "[](){}", "[]()[]", "[]{()}", "[]{{}}", "[]{[]}", "[]{}()", "[]{}{}", "[]{}[]", "[][()]", "[][{}]", "[][[]]", "[][]()", "[][]{}", "[][][]"}},
	} {
		if got := GenBrackets(tc.input); !reflect.DeepEqual(tc.want, got) {
			t.Errorf("CorrectParentheses(%v) = %v, want = %v", tc.input, got, tc.want)
		}
	}
}