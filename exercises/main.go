package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	opening = "([{"
	closing = ")]}"
)

func CorrectParentheses(s string) bool {
	var counter int
	var correct func(s string) (bool, int, string)
	correct = func(s string) (bool, int, string) {
		if s == "" {
			return counter == 0, -1, ""
		}
		r, rn := utf8.DecodeRuneInString(s) // Getting first rune in string
		s = s[rn:]                          // Consuming element
		if j := strings.IndexRune(opening, r); j != -1 {
			counter++
			valid, k, ns := correct(s)
			if !valid || j != k {
				return false, -1, ""
			}
			return correct(ns)
		} else if j := strings.IndexRune(closing, r); j != -1 {
			counter--
			return true, j, s
		}

		return false, 0, s // Should never happen
	}

	result, _, _ := correct(s)
	return result
}

// GenParentheses generates all valid parentheses pair of length n*2.
func GenParentheses(n int, f func(string)) {
	var gen func(o, c int, s string)
	gen = func(o, c int, s string) {
		if o+c == 2*n {
			if CorrectParentheses(s) {
				f(s)
			}
			return
		}
		if o < n {
			gen(o+1, c, s+"(")
			gen(o+1, c, s+"[")
			gen(o+1, c, s+"{")
		}
		if o > c {
			gen(o, c+1, s+")")
			gen(o, c+1, s+"]")
			gen(o, c+1, s+"}")
		}
	}
	gen(0, 0, "")
}

func main() {
	input := "([{}]))"
	fmt.Println(CorrectParentheses(input))
}
