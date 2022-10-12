package main

const (
	openParentheses  = "([{"
	closeParentheses = ")]}"
)

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
		if len(opened) > 0 {
			if j := opened[len(opened)-1]; o > c {
				gen(o, c+1, s+string(closeParentheses[j]), opened[:len(opened)-1])
			}
		}
	}
	gen(0, 0, "", nil)
}
