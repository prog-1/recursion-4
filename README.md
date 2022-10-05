# Parentheses Generation (Recursion Part IV)

## Correctness (single)

### Examples

```text
Incorrect
(()())(
)()(

Correct
(()())()
()()()
```

### Code

https://go.dev/play/p/TVUd0p006Hz

```go
// CorrectParentheses checks whether parentheses are correct in s.
func CorrectParentheses(s string) bool {
	opened := 0
	for _, c := range s {
		switch c {
		case '(':
			opened++
		case ')':
			opened--
		}
		if opened < 0 {
			return false
		}
	}
	return opened == 0
}
```

### Exercise

Implement a recursive version of `CorrectParentheses`.

## Correctness (many)

### Examples

```text
Incorrect
({[}])
(}[])

Correct
()[]{}
([{}])
```

### Code

https://go.dev/play/p/j6jWlQ_IkTw

```go
// CorrectParentheses checks whether parentheses are correct in s.
func CorrectParentheses(s string) bool {
	var opened []int // Stack, last in - first out (LIFO).
	for _, c := range s {
		if i := strings.IndexRune(openParentheses, c); i >= 0 {
			opened = append(opened, i)
		} else if i := strings.IndexRune(closeParentheses, c); i >= 0 {
			if len(opened) == 0 {
				return false
			}
			var j int
			j, opened = opened[len(opened)-1], opened[:len(opened)-1] // Slice tricks: pop.
			if i != j {
				return false
			}
		}
	}
	return len(opened) == 0
}
```

### Exercise

Implement a recursive version of `CorrectParentheses`.

## Recursive Generator

### Code

https://go.dev/play/p/Eq-8crxAeuc

```go
// GenParentheses generates all valid parentheses pair of length n*2.
func GenParentheses(n int, f func(string)) {
	var gen func(o, c int, s string)
	gen = func(o, c int, s string) {
		if o+c == 2*n {
			f(s)
			return
		}
		if o < n {
			gen(o+1, c, s+"(")
		}
		if o > c {
			gen(o, c+1, s+")")
		}
	}
	gen(0, 0, "")
}
```

### Exercise

Implement recursive `GenParentheses` for `()[]{}`.
