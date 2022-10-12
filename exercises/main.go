package main

import (
	"fmt"
	"strings"
)

const (
	opening = "([{"
	closing = ")]}"
)

// // CorrectParentheses checks whether parentheses are correct in s.
// func CorrectParenthesesStack(s string) bool {
// 	var Correct func(s string, i int, o []int) bool
// 	Correct = func(s string, i int, o []int) bool {
// 		if i > len(s)-1 { // Exit condition
// 			return len(o) == 0
// 		}
// 		r := rune(s[i])
// 		if j := strings.IndexRune(opening, r); j != -1 { // Opened
// 			o = append(o, j)
// 		}
// 		if j := strings.IndexRune(closing, r); j != -1 { // Closed
// 			if len(o) == 0 { // No opened
// 				return false
// 			}
// 			var lo int
// 			lo, o = o[len(o)-1], o[:len(o)-1] // Pop
// 			if lo != j {
// 				return false
// 			}
// 		}
// 		return Correct(s, i+1, o)
// 	}
// 	return Correct(s, 0, nil)
// }

func CorrectParentheses(s string) bool {
	var counter int
	var correct func(s string, i int, counter *int) (bool, int, int)
	correct = func(s string, i int, counter *int) (bool, int, int) {
		if i > len(s)-1 {
			return *counter == 0, -1, -1
		}
		r := rune(s[i])
		if j := strings.IndexRune(opening, r); j != -1 {
			*counter++
			_, a, b := correct(s, i+1, counter)
			if j != a {
				return false, -1, -1
			}
			return correct(s, b, counter)
		}
		if j := strings.IndexRune(closing, r); j != -1 {
			*counter--
			return false, j, i + 1
		}

		return false, 0, 0
	}

	result, _, _ := correct(s, 0, &counter)
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
	var result []string
	GenParentheses(3, func(s string) { result = append(result, s) })
	fmt.Println(result)
}
