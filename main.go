package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	errorCode        = math.MinInt
	openParentheses  = "([{"
	closeParentheses = ")]}"
)

func index(r rune) int {
	tmp := strings.IndexRune(openParentheses, r)
	if tmp != -1 {
		return tmp + 1
	}
	return (strings.IndexRune(closeParentheses, r) + 1) * -1
}

func CorrectParentheses(s string) bool {
	got := check("(" + s + ")")
	return len(s)+2 == got
}

func check(s string) int {
	symbol := index(rune(s[0]))
	if len(s) <= 1 {
		return errorCode
	}
	if symbol < 0 {
		return errorCode
	}
	next := index(rune(s[1]))
	if next < 0 {
		if next != symbol*-1 {
			return errorCode
		}
		return 2
	} else {
		var i int
		for {
			if index(rune(s[i+1])) == symbol*-1 {
				return i + 2
			}
			tmp := check(s[i+1:])
			if tmp == errorCode {
				return errorCode
			}
			if len(s)-1 <= tmp+i {
				return errorCode
			}
			i += tmp
		}
	}

}

func main() {
	fmt.Println(CorrectParentheses("{({})}{}"))
}
