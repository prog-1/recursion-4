package main

import "fmt"

// GenParentheses generates all valid parentheses pair of length n*2.

func GenParentheses(n int, f func(string)) {
	op := []string{"(", "[", "{"}
	cl := []string{")", "]", "}"}
	var gen func(o, c int, s string, x []int)
	gen = func(o, c int, s string, x []int) {
		if o+c == 2*n {
			f(s)
			return
		}
		for i := range op {
			if o < n {
				gen(o+1, c, s+string(op[i]), append(x, i))
			}
		}
		if len(x) > 0 {
			if j := x[len(x)-1]; o > c {
				gen(o, c+1, s+cl[j], x[:len(x)-1])
			}
		}
	}
	gen(0, 0, "", nil)
}

func main() {
	GenParentheses(3, func(s string) { fmt.Println(s) })
}
