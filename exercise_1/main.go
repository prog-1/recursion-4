package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	openParentheses  = "([{"
	closeParentheses = ")]}"
)

// CorrectParentheses checks whether parentheses are correct in s.
func CorrectParentheses(s string) (ok bool) {
	for s != "" {
		if s, ok = CheckParentheses(s); !ok {
			return false
		}
	}
	return true
}

func CheckParentheses(s string) (string, bool) {
	if s == "" {
		return "", true
	}
	c, cn := utf8.DecodeRuneInString(s)
	if cn <= 0 { // Unable to decode
		return s, false
	}
	if strings.ContainsRune(closeParentheses, c) { // First cannot be closing
		return s, false
	}
	i := strings.IndexRune(openParentheses, c)
	if i >= 0 { // Found opening
		s = s[cn:]
	} else {
		CheckParentheses(s[cn:]) // Continue checking
	}
	cnt := 1
	for {
		cc, cn := utf8.DecodeRuneInString(s)
		if cn <= 0 { // Unable to decode
			return s, false
		}
		if cc == c { // The same opening
			s = s[cn:]
			cnt++
		} else if j := strings.IndexRune(closeParentheses, cc); j == i { // Corresponding closing
			s = s[cn:]
			cnt--
			if cnt == 0 { // Closed all previously open
				return s, true
			}
		} else { // Probably another type of opening
			var ok bool
			if s, ok = CheckParentheses(s); !ok {
				return s, ok
			}
		}
	}
}

func main() {
	fmt.Println(CorrectParentheses("()[]{}")) // true
	fmt.Println(CorrectParentheses("([{}])")) // true
	fmt.Println(CorrectParentheses("({[}])")) // false
}
