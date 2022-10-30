package main

import "fmt"

const (
	openParentheses  = "([{"
	closeParentheses = ")]}"
)

// GenParentheses generates all valid parentheses pair of length n*2.
func GenParentheses(n int, f func(string)) {
	var gen func(o, c int, s string, opened []int)
	gen = func(o, c int, s string, opened []int) {
		if o+c == 2*n {
			f(s)
			return
		}
		for i := 0; i < len(openParentheses); i++ {
			if o < n {
				gen(o+1, c, s+string(openParentheses[i]), append(opened, i))
			}
		}
		if len(opened) > 0 && o > c {
			var j int
			j, opened = opened[len(opened)-1], opened[:len(opened)-1] // Slice tricks: pop.
			gen(o, c+1, s+string(closeParentheses[j]), opened)
		}
	}
	gen(0, 0, "", nil)
}

func main() {
	GenParentheses(2, func(s string) { fmt.Println(s) })
}
