package main

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

// func main() {
// 	GenParentheses(3, func(s string) { fmt.Println(s) })
// }
