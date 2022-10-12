package main

import (
	"fmt"
	"strings"
)

const (
	open  = "([{"
	close = ")]}"
)

func CorrectParentheses(s string) bool {
	var cnt int
	var cor func(s string, i int, cnt *int) (bool, int, int)
	cor = func(s string, i int, cnt *int) (bool, int, int) {
		if i > len(s)-1 {
			return *cnt == 0, -1, -1
		}
		r := rune(s[i])
		if x := strings.IndexRune(open, r); x > -1 {
			*cnt++
			_, z, c := cor(s, i+1, cnt)
			if x != z {
				return false, -1, -1
			}
			return cor(s, c, cnt)
		}
		if x := strings.IndexRune(close, r); x > -1 {
			*cnt--
			return false, x, i + 1
		}
		return false, 0, 0
	}
	res, _, _ := cor(s, 0, &cnt)
	return res
}

func main() {
	for _, s := range []string{
		"(()())(",
		")()(",
		"({[}])",
		"(}[])",
		"(()())()",
		"()()()",
		"()[]{}",
		"([{}])",
	} {
		fmt.Println(CorrectParentheses(s), s)
	}
}
