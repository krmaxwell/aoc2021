package day10

type RuneStack struct {
	items []rune
	top   int
}

func ValidateChunks(s string) int {
	seen := RuneStack{}
	expected := RuneStack{}
	for _, c := range s {
		switch c {
		case '(':
			seen.Push(c)
			expected.Push(')')
		case '{':
			seen.Push(c)
			expected.Push('}')
		case '<':
			seen.Push(c)
			expected.Push('>')
		case '[':
			seen.Push(c)
			expected.Push(']')
		case ']', '>', '}', ')':
			if expected.Peek() == c {
				// pop from both stacks, continues to be valid
				expected.Pop()
				seen.Pop()
			} else {
				switch c {
				case ')':
					return 3
				case ']':
					return 57
				case '}':
					return 1197
				case '>':
					return 25137
				}
			}
		}
	}
	return 0
}

func (r *RuneStack) Push(c rune) {
	r.items = append(r.items, c)
	r.top++
}

func (r *RuneStack) Peek() rune {
	c := r.items[r.top-1]
	return c
}

func (r *RuneStack) Pop() rune {
	c := r.items[r.top-1]
	r.items = r.items[:r.top-1]
	r.top--
	return c
}
