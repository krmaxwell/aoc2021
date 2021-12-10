package day10

func TestLine(s string) int {
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '{' || c == '<' || c == '[' {
			stack = append(stack, c) // push
		} else if c == ')' {
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1] // pop
			} else {
				return 3
			}
		} else if c == '}' {
			if stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1] // pop
			} else {
				return 1197
			}
		} else if c == '>' {
			if stack[len(stack)-1] == '<' {
				stack = stack[:len(stack)-1] // pop
			} else {
				return 25137
			}
		} else if c == ']' {
			if stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1] // pop
			} else {
				return 57
			}
		}
	}
	return 0
}
