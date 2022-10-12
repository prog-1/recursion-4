package main

import "fmt"

//not really my solution
//https://www.geeksforgeeks.org/check-for-balanced-parenthesis-without-using-stack/

func closeBreckets(c byte) byte {
	switch c {
	case '(':
		return ')'
	case '{':
		return '}'
	case '[':
		return ']'
	}
	return ' '
}

func correctParentheses(s string, n int) bool {
	if n == 0 {
		return true
	}
	if n == 1 || s[0] == ')' || s[0] == '}' || s[0] == ']' {
		return false
	}

	i, count := 0, 0
	for i = 1; i < n; i++ {
		if s[i] == s[0] {
			count++
		}
		if s[i] == closeBreckets(s[0]) {
			if count == 0 {
				break
			}
			count--
		}
	}
	if i == n {
		return false
	}
	return correctParentheses(s[1:], i-1) && correctParentheses(s[i+1:], n-i-1)
}

func main() {
	s := string("([{}]{]})")
	s2 := string("()([]))")
	n := len(s)
	n2 := len(s2)
	fmt.Println(correctParentheses(s, n))
	fmt.Println(correctParentheses(s2, n2))
}
