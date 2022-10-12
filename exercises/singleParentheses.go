package main

// Algorithm:
// var counter int
// We parse string and get element
// If parentheses is an open one, then increase the counter by 1
// If it's closed one, then decrease the counter by 1
// Check if counter > 0, because we can't ever have more closing brackets, than open ones
// After check whether we have the same amount of open and closed brackets.

func CorrectParenthesesLoop(s string) bool {
	c := 0

	for _, el := range s {
		if el == '(' {
			c++
		} else if el == ')' {
			c--
		}
		if c < 0 {
			return false // We can't have more closing parentheses than open ones
		}
	}

	return c == 0 // We can't have unequal amount of brackets
}

func CorrectParenthesesRecursive(s string) bool {
	c := 0
	var correct func(s string, i int, c *int) bool
	correct = func(s string, i int, c *int) bool {

		if i > len(s)-1 {
			return *c == 0
		}

		if s[i] == '(' {
			*c++
		} else if s[i] == ')' {
			*c--
		}
		if *c < 0 {
			return false // We can't have more closing parentheses than open ones
		}

		return correct(s, i+1, c)
	}
	return correct(s, 0, &c)
}

// func main() {
// 	s := "()()())"
// 	fmt.Println(CorrectParenthesesRecursive(s))
// }
