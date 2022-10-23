package main

// CorrectParentheses checks whether parentheses are correct in s.
func CorrectParentheses(s string) bool {
	sl := make([]string, 0)
	for i := 0; i < len(s); i++ {
		sl = append(sl, string(s[i]))
	}
	var cor func(sl []string) ([]string, bool)
	cor = func(sl []string) ([]string, bool) {
		if len(sl) == 0 {
			return sl, true
		}
		if sl[0] == "(" || sl[0] == "[" || sl[0] == "{" {
			for i := 0; i < len(sl); i++ {
				if sl[i] == ")" {
					sl = append(sl[:i], sl[i+1:]...)
					sl = sl[1:]
					return cor(sl)
				}
				if sl[i] == "]" {
					sl = append(sl[:i], sl[i+1:]...)
					sl = sl[1:]
					return cor(sl)
				}
				if sl[i] == "}" {
					sl = append(sl[:i], sl[i+1:]...)
					sl = sl[1:]
					return cor(sl)
				}
			}
		} else {
			return sl, false
		}
		return sl, false
	}
	_, res := cor(sl)
	return res

}

// func main() {
// 	for _, s := range []string{
// 		"(()())(",
// 		")()(",
// 		"({[}])",
// 		"(}[])",
// 		"(()())()",
// 		"()()()",
// 		"()[]{}",
// 		"([{}])",
// 	} {
// 		fmt.Println(CorrectParentheses(s), s)
// 	}
// }
