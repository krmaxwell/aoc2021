package day10

type RuneStack struct {
	items []rune
	top   int
}

type SubsystemResult struct {
	Score int
	Valid bool
}

func ValidateChunks(s string) SubsystemResult {
	seen := RuneStack{}
	expected := RuneStack{}
	result := SubsystemResult{Score: 0, Valid: false}
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
					result.Score = 3
					return result
				case ']':
					result.Score = 57
					return result
				case '}':
					result.Score = 1197
					return result
				case '>':
					result.Score = 25137
					return result
				}
			}
		}
	}
	result.Valid = true
	result.Score = CompletionScore(expected)
	return result
}

func CompletionScore(r RuneStack) int {
	total := 0
	for i := len(r.items); i > 0; i-- {
		c := r.Pop()
		total *= 5
		switch c {
		case ')':
			total += 1
		case ']':
			total += 2
		case '}':
			total += 3
		case '>':
			total += 4
		}
	}
	return total
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
