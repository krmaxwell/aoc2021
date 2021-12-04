package day4

type BoardLayout [5][5]int
type BoardState [5][5]bool
type Board struct {
	Layout BoardLayout
	State  BoardState
	Round  int
	HasWon bool
	Score  int
}

func (b Board) CalcResult(draw []int) Board {
	for round, pick := range draw {
		found := false
		for row := range b.Layout { // replace with range
			for col := range b.Layout[row] {
				if b.Layout[row][col] == pick {
					b.State[row][col] = true
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		b.Round = round + 1 // because of 0 index
		if b.State.won() {
			b.HasWon = true
			b.Score = b.CalcScore(pick)
			return b
		}
	}

	return b
}

func (b Board) CalcScore(pick int) int {
	score := 0
	for row := range b.Layout {
		for col := range b.Layout[row] {
			if !b.State[row][col] {
				score += b.Layout[row][col]
			}
		}
	}
	score *= pick
	return score
}

func (s BoardState) won() bool {
	result := false
	for _, row := range s {
		result = checkLine(row)
		if result {
			return result
		}
	}
	for colNum := range s {
		var column [5]bool
		for i := 0; i < 5; i++ {
			column[i] = s[i][colNum]
		}
		result = checkLine(column)
		if result {
			return result
		}
	}
	return result
}

func checkLine(line [5]bool) bool {
	return line[0] && line[1] && line[2] && line[3] && line[4]
}
