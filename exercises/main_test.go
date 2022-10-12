package main

func TestCheck(t *testing.T) {
	for t := range []struct{
		input 
		want bool
	} {
		{"([{}])", true},
		{"({[}])", false},
	} {
		
	}
}